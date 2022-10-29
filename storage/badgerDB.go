package storage

import (
	"context"
	"encoding/json"
	"github.com/dgraph-io/badger/v3"
	"github.com/ljg-cqu/biance/logger"
	"github.com/ljg-cqu/biance/utils/backoff"
	"github.com/pkg/errors"
	"time"
)

var _ DB = (*BadgerDB)(nil)

type BadgerDB struct {
	logger.Logger
	*badger.DB
}

func New(path string, logger logger.Logger, loggerAdapter logger.BadgerDBLogger) *BadgerDB {
	opt := badger.DefaultOptions(path)
	opt.Logger = loggerAdapter
	db, err := badger.Open(opt)
	if err != nil {
		panic("Failed to open badger database, error:" + err.Error())
	}
	return &BadgerDB{logger, db}
}

// Classic k-v storage

func (b *BadgerDB) Set(ctx context.Context, key, val []byte) (err error) {
	err = backoff.RetryFnExponentialForever(b.Logger, ctx, time.Second, time.Second*10, func() (bool, error) {
		err = b.Update(func(txn *badger.Txn) error {
			err := txn.Set(key, val)
			return errors.WithStack(err)
		})
		if err != nil {
			return true, errors.WithStack(err)
		}
		return false, nil
	})
	err = errors.Wrapf(err, "failed to set k-v. key:%v, val:%v", string(key), string(val))
	return
}

func (b *BadgerDB) Get(ctx context.Context, key []byte) (value []byte, err error) {
	err = backoff.RetryFnExponential10Times(b.Logger, ctx, time.Second, time.Second*10, func() (bool, error) {
		err = b.View(func(txn *badger.Txn) error {
			item, err := txn.Get(key)
			if err != nil {
				return errors.WithStack(err)
			}

			err = item.Value(func(val []byte) error {
				value = val
				return nil
			})
			return errors.WithStack(err)
		})
		if err != nil {
			if errors.Is(err, badger.ErrKeyNotFound) {
				return false, errors.WithStack(err)
			}
			return true, errors.WithStack(err)
		}
		return false, nil
	})
	err = errors.Wrapf(err, "failed to get value by key. key:%v", string(key))
	return
}

func (b *BadgerDB) List(ctx context.Context, prefix []byte) ([][]byte, error) {
	var valBytesArr [][]byte
	err := backoff.RetryFnExponential10Times(b.Logger, ctx, time.Second, time.Second*10, func() (bool, error) {
		err := b.View(func(txn *badger.Txn) error {
			it := txn.NewIterator(badger.DefaultIteratorOptions)
			defer it.Close()
			for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
				item := it.Item()
				err := item.Value(func(v []byte) error {
					var valBytes = make([]byte, len(v))
					copy(valBytes, v)
					valBytesArr = append(valBytesArr, valBytes)
					return nil
				})
				if err != nil {
					return errors.WithStack(err)
				}
			}
			return nil
		})
		if err != nil {
			if errors.Is(err, badger.ErrKeyNotFound) {
				return false, errors.WithStack(err)
			}
			return true, errors.WithStack(err)
		}
		return false, nil
	})
	err = errors.Wrapf(err, "failed to list values by prefix. prefix:%v", string(prefix))
	return valBytesArr, nil
}

// With marshal and unmarshal support

func (b *BadgerDB) MSet(ctx context.Context, key []byte, val interface{}) error {
	valBytes, err := json.Marshal(val)
	if err != nil {
		return errors.WithStack(err)
	}

	err = b.Set(ctx, key, valBytes)

	return errors.WithStack(err)
}

func (b *BadgerDB) MGet(ctx context.Context, key []byte, val interface{}) error {
	valBytes, err := b.Get(ctx, key)
	if err != nil {
		return errors.WithStack(err)
	}

	err = json.Unmarshal(valBytes, val)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (b *BadgerDB) MList(ctx context.Context, prefix []byte, val interface{}) error {
	return errors.New("to to implemented") // todo
}

// With Model(s) interface support

func (b *BadgerDB) SaveModel(ctx context.Context, data Model) error {
	return errors.WithStack(b.MSet(ctx, data.Key(), data))
}

func (b *BadgerDB) LoadModel(ctx context.Context, data Model) error {
	return errors.WithStack(b.MGet(ctx, data.Key(), data))
}

func (b *BadgerDB) ListModels(ctx context.Context, datum Models) error {
	return errors.WithStack(b.MList(ctx, datum.Prefix(), datum))
}
