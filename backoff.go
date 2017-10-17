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
	return time.Duration(rand.Float64()*(1<<attempt)) * time.Second
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
	return time.Duration(rand.Float64()*(1<<n)) * time.Second
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
