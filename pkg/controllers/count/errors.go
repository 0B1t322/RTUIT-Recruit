package count

import (
	"errors"
)

var (
	ErrNotFound = errors.New("Count not found")
	ErrExist 	= errors.New("Count exist")
)