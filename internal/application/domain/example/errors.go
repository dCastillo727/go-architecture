package example

import "errors"

var (
	ErrExampleNotFound = errors.New("The requested example was not found")
)
