package template

import (
	"time"

	"github.com/go-kit/kit/log"
)

// loggingMiddleware -
type loggingMiddleware struct {
	logger log.Logger
	next   Service
}

// LoggingMiddleware -
func LoggingMiddleware(logger log.Logger) ServiceMiddleware {

	return func(svc Service) Service {
		return loggingMiddleware{
			logger: logger,
			next:   svc,
		}
	}
}

// Template -
func (mw loggingMiddleware) Template(req Request) (resp Response, err error) {

	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "template",
			"input", req.Test,
			"output", resp.Test,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	return mw.next.Template(req)
}
