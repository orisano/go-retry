package retry

import (
	"time"

	"github.com/pkg/errors"
)

type Temporary interface {
	Temporary() bool
}

type After interface {
	RetryAfter() (bool, time.Duration)
}

func isTemporary(err error) bool {
	te, ok := errors.Cause(err).(Temporary)
	return ok && te.Temporary()
}

func Do(do func() error, opts ...option) error {
	options := defaultOption()
	for _, opt := range opts {
		opt(options)
	}

	attempt := uint32(0)
	for {
		err := do()
		if err == nil {
			return nil
		}

		attempt++
		if attempt >= options.MaxAttempt {
			return errors.New("max attempt exceeded")
		}
		if ra, ok := err.(After); ok {
			wait, d := ra.RetryAfter()
			if wait {
				time.Sleep(d)
				continue
			}
		}
		if options.ForceRetry || isTemporary(err) {
			d := options.RetryPolicy.Backoff(attempt)
			time.Sleep(d)
		} else {
			return err
		}
	}
	return nil
}
