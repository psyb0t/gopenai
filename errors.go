package gopenai

import "errors"

var (
	// ErrRequestTimeout is an error that indicates a request has timed out.
	ErrRequestTimeout = errors.New("request timeout")
)
