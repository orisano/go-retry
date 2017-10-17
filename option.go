package retry

type options struct {
	MaxAttempt  uint32
	ForceRetry  bool
	RetryPolicy Backoff
}
type option func(*options)

func defaultOption() *options {
	return &options{
		MaxAttempt:  10,
		ForceRetry:  false,
		RetryPolicy: TruncatedExponentialBackoff(10),
	}
}

func MaxAttempt(maxAttempt uint32) option {
	return func(options *options) {
		options.MaxAttempt = maxAttempt
	}
}

func Force() option {
	return func(options *options) {
		options.ForceRetry = true
	}
}

func WithPolicy(backoff Backoff) option {
	return func(options *options) {
		options.RetryPolicy = backoff
	}
}
