package retry

import (
	"math/rand"
	"time"
)

type Backoff interface {
	Backoff(attempt uint32) time.Duration
}

type exponentialBackoff struct{}

func ExponentialBackoff() Backoff {
	return &exponentialBackoff{}
}

func (*exponentialBackoff) Backoff(attempt uint32) time.Duration {
	return time.Duration(rand.Int63n((1 << attempt) * int64(time.Second)))
}

type truncatedExponentialBackoff struct {
	max uint32
}

func TruncatedExponentialBackoff(max uint32) Backoff {
	return &truncatedExponentialBackoff{max}
}

func (b *truncatedExponentialBackoff) Backoff(attempt uint32) time.Duration {
	n := attempt
	if n > b.max {
		n = b.max
	}
	return time.Duration(rand.Int63n((1 << n) * int64(time.Second)))
}

type constantBackoff struct {
	duration time.Duration
}

func ConstantBackoff(d time.Duration) Backoff {
	return &constantBackoff{d}
}

func (b *constantBackoff) Backoff(uint32) time.Duration {
	return b.duration
}
