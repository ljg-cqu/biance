package email

import (
	"context"
	"crypto/tls"
	"github.com/ljg-cqu/biance/logger"
	"github.com/ljg-cqu/biance/utils/backoff"
	"github.com/ljg-cqu/core/smtp"
	"github.com/pkg/errors"
	mail "github.com/xhit/go-simple-mail/v2"
	"time"
)

func SendPNLReportWith126Mail(log logger.Logger, ctx context.Context, subject, content string) error {
	email := mail.NewMSG()
	email.SetFrom("Zealy <ljg_cqu@126.com>").
		AddTo("ljg_cqu@163.com").
		SetSubject(subject)
	email.SetBody(mail.TextPlain, content)
	err := backoff.RetryFnExponential10Times(log, ctx, time.Second, time.Second*10, func() (bool, error) {
		emailCli, err := smtp.NewEmailClient(smtp.NetEase126Mail, &tls.Config{InsecureSkipVerify: true}, "ljg_cqu@126.com", "XROTXFGWZUILANPB")
		if err != nil {
			return false, errors.Wrapf(err, "failed to create 126 email client.")
		}
		err = emailCli.Send(email)
		if err != nil {
			return false, errors.Wrapf(err, "failed to send 126 email")
		}
		return false, nil
	})
	return errors.WithStack(err)
}

func SendPNLReportWith163Mail(log logger.Logger, ctx context.Context, subject, content string) error {
	email := mail.NewMSG()
	email.SetFrom("Zealy <ljg_cqu@163.com>").
		AddTo("ljg_cqu@163.com").
		SetSubject(subject)
	email.SetBody(mail.TextPlain, content)

	err := backoff.RetryFnExponential10Times(log, ctx, time.Second, time.Second*10, func() (bool, error) {
		emailCli, err := smtp.NewEmailClient(smtp.NetEase163Mail, &tls.Config{InsecureSkipVerify: true}, "ljg_cqu@163.com", "QLVMFGJIHJEPBZOR")
		if err != nil {
			return false, errors.Wrapf(err, "failed to create 163 email client.")
		}
		err = emailCli.Send(email)
		if err != nil {
			return false, errors.Wrapf(err, "failed to send 163 email")
		}
		return false, nil
	})
	return errors.WithStack(err)
}

func SendPNLReportWithQQMail(log logger.Logger, ctx context.Context, subject, content string) error {
	email := mail.NewMSG()
	email.SetFrom("Zealy <1025003548@qq.com>").
		AddTo("ljg_cqu@163.com").
		SetSubject(subject)
	email.SetBody(mail.TextPlain, content)

	err := backoff.RetryFnExponential10Times(log, ctx, time.Second, time.Second*10, func() (bool, error) {
		emailCli, err := smtp.NewEmailClient(smtp.QQMail, &tls.Config{InsecureSkipVerify: true}, "1025003548@qq.com", "ncoajiivbenpbfbh")
		if err != nil {
			return false, errors.Wrapf(err, "failed to create QQ email client.")
		}
		err = emailCli.Send(email)
		if err != nil {
			return false, errors.Wrapf(err, "failed to send QQ email")
		}
		return false, nil
	})
	return errors.WithStack(err)
}
