package shop

import (
	"errors"
)

var (
	ErrNotFound  		= errors.New("Shop not found")
	ErrNegCount			= errors.New("Count can't be negative")
	ErrProductNotFound	= errors.New("Product not found")
)