package fromqq126163

import (
	"context"
	"github.com/ljg-cqu/biance/email"
	"github.com/ljg-cqu/biance/logger"
)

func Send(logger logger.Logger, ctx context.Context, subject, content, to string) {
	err := email.SendPNLReportWith163Mail(logger, ctx, subject, content, to)
	logger.DebugOnError(err, "Failed to send email with 163 mailbox")

	if err != nil {
		err = email.SendPNLReportWith126Mail(logger, ctx, subject, content, to)
		logger.DebugOnError(err, "Failed to send email with 126 mailbox")
	}

	if err != nil {
		err := email.SendPNLReportWithQQMail(logger, ctx, subject, content, to)
		logger.DebugOnError(err, "Failed to send email with QQ mailbox.")
	}
}
