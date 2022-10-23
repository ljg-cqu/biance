package backoff

import (
	"context"
	"github.com/ljg-cqu/logger"
	"github.com/lestrrat-go/backoff/v2"
	"github.com/pkg/errors"
	"time"
)

// Retry the retry function, with outer constant policy and inner exponential policy

func RetryRetryFnForever(log logger.Logger, ctx context.Context, fn Fn) error {
	return RetryRetryFn(log, ctx, ConstantPolicyForever(time.Minute*10), ExponentialPolicy10Times(time.Second, time.Second*10), fn)
}

func RetryRetryFn100Times(log logger.Logger, ctx context.Context, fn Fn) error {
	return RetryRetryFn(log, ctx, ConstantPolicy(100, time.Minute*10), ExponentialPolicy10Times(time.Second, time.Second*10), fn)
}

func RetryRetryFn10Times(log logger.Logger, ctx context.Context, fn Fn) error {
	return RetryRetryFn(log, ctx, ConstantPolicy(10, time.Minute*10), ExponentialPolicy10Times(time.Second, time.Second*10), fn)
}

func RetryRetryFn(log logger.Logger, ctx context.Context, policyOuter backoff.Policy, policyInner backoff.Policy, fn Fn) error {
	err := RetryFn(log, ctx, policyOuter, func() (retry bool, err error) {
		if err := RetryFn(log, ctx, policyInner, fn); err != nil {
			return true, errors.Wrapf(err, "inner retry error")
		}
		return false, nil
	})
	return errors.Wrapf(err, "outter retry error")
}

// Retry with constant policy

func RetryFnConstantForever(log logger.Logger, ctx context.Context, dur time.Duration, fn Fn) error { // handy function
	policy := ConstantPolicy(0, dur)
	return RetryFn(log, ctx, policy, fn)
}

func RetryFnConstant100Times(log logger.Logger, ctx context.Context, dur time.Duration, fn Fn) error {
	policy := ConstantPolicy(100, dur)
	return RetryFn(log, ctx, policy, fn)
}

func RetryFnConstant30Times(log logger.Logger, ctx context.Context, dur time.Duration, fn Fn) error {
	policy := ConstantPolicy(30, dur)
	return RetryFn(log, ctx, policy, fn)
}

func RetryFnConstant10Times(log logger.Logger, ctx context.Context, dur time.Duration, fn Fn) error {
	policy := ConstantPolicy(10, dur)
	return RetryFn(log, ctx, policy, fn)
}

func RetryFnConstant(log logger.Logger, ctx context.Context, maxRetries int, dur time.Duration, fn Fn) error {
	policy := ConstantPolicy(maxRetries, dur)
	return RetryFn(log, ctx, policy, fn)
}

// Retry with exponential policy

func RetryFnExponentialForever(log logger.Logger, ctx context.Context, minInterval, maxInterval time.Duration, fn Fn) error { // handy function
	policy := ExponentialPolicy(0, minInterval, maxInterval)
	return RetryFn(log, ctx, policy, fn)
}

func RetryFnExponential100Times(log logger.Logger, ctx context.Context, minInterval, maxInterval time.Duration, fn Fn) error {
	policy := ExponentialPolicy(100, minInterval, maxInterval)
	return RetryFn(log, ctx, policy, fn)
}

func RetryFnExponential30Times(log logger.Logger, ctx context.Context, minInterval, maxInterval time.Duration, fn Fn) error {
	policy := ExponentialPolicy(30, minInterval, maxInterval)
	return RetryFn(log, ctx, policy, fn)
}

func RetryFnExponential10Times(log logger.Logger, ctx context.Context, minInterval, maxInterval time.Duration, fn Fn) error {
	policy := ExponentialPolicy(10, minInterval, maxInterval)
	return RetryFn(log, ctx, policy, fn)
}

func RetryFnExponential(log logger.Logger, ctx context.Context, maxRetries int, minInterval, maxInterval time.Duration, fn Fn) error {
	policy := ExponentialPolicy(maxRetries, minInterval, maxInterval)
	return RetryFn(log, ctx, policy, fn)
}

// Constant policy

func ConstantPolicyForever(dur time.Duration) backoff.Policy {
	return ConstantPolicy(0, dur)
}

func ConstantPolicy100Times(dur time.Duration) backoff.Policy {
	return ConstantPolicy(100, dur)
}

func ConstantPolicy30Times(dur time.Duration) backoff.Policy {
	return ConstantPolicy(30, dur)
}

func ConstantPolicy10Times(dur time.Duration) backoff.Policy {
	return ConstantPolicy(10, dur)
}

func ConstantPolicy(maxRetries int, dur time.Duration) backoff.Policy {
	p := backoff.Constant(
		backoff.WithMaxRetries(maxRetries), // 0 for forever retry
		backoff.WithInterval(dur))
	return p
}

// Exponential policy

func ExponentialPolicyForever(minInterval, maxInterval time.Duration) backoff.Policy {
	return ExponentialPolicy(0, minInterval, maxInterval)
}

func ExponentialPolicy100Times(minInterval, maxInterval time.Duration) backoff.Policy {
	return ExponentialPolicy(100, minInterval, maxInterval)
}

func ExponentialPolicy30Times(minInterval, maxInterval time.Duration) backoff.Policy {
	return ExponentialPolicy(30, minInterval, maxInterval)
}

func ExponentialPolicy10Times(minInterval, maxInterval time.Duration) backoff.Policy {
	return ExponentialPolicy(10, minInterval, maxInterval)
}

func ExponentialPolicy(maxRetries int, minInterval, maxInterval time.Duration) backoff.Policy {
	p := backoff.Exponential(
		backoff.WithMaxRetries(maxRetries), // 0 for forever retry
		backoff.WithMinInterval(minInterval),
		backoff.WithMaxInterval(maxInterval),
		backoff.WithMultiplier(1.5),
		backoff.WithJitterFactor(0.05),
	)
	return p
}

// The basic retry function

type Fn func() (retry bool, err error)

func RetryFn(log logger.Logger, ctx context.Context, policy backoff.Policy, fn Fn) error {
	b := policy.Start(ctx)
	var errStack error
	var startAt = time.Now()
	var retryNum = 0
	var lastRetryAt = time.Now()
	for backoff.Continue(b) {
		retry, err := fn()
		errStack = errors.WithStack(err)
		if !retry {
			break
		}
		retryNum++
		log.Error("Retry", []logger.Field{
			{"error", err},
			{"retryNum", retryNum},
			{"retryAfter", time.Now().Sub(lastRetryAt).Seconds()}}...)
		lastRetryAt = time.Now()
	}
	retryDur := time.Now().Sub(startAt).Seconds()
	return errors.Wrapf(errStack, "exited after retrying %v times in %v seconds", retryNum, retryDur)
}
