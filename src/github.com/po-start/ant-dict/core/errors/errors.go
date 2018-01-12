package errors

import (
	"errors"
)

var (
	ErrFileAccessFailure = errors.New("file access failure")
	ErrAddressNull       = errors.New("address is nil")
)
