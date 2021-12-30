package errs

import "errors"

var (
	ErrNotFound       = errors.New("not found")
	ErrInternal       = errors.New("something went wrong")
	ErrAccessDenied   = errors.New("access denied")
	ErrInvalidRequest = errors.New("invalid request")
	ErrEmailTaken     = errors.New("this email already taken")
	ErrUsernameTaken  = errors.New("this username already taken")
)
