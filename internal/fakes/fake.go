package fakes

import "errors"

var notImplementedError = errors.New("not implemented")

type fakeService struct {
	Calls []CallRecord

	// Err is the error returned with every implemented method.
	// Useful for faking client errors.
	Err error
}
