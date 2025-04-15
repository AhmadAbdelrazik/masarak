package domain

import "errors"

var (
	ErrInvalidProperty = errors.New("invalid property")
	ErrInvalidUpdate   = errors.New("invalid update")
)
