package repository

import "errors"

var (
	ErrObjectExists = errors.New("Object with this ID already exists")
	ErrUknownId     = errors.New("Id doesn't exist")
)
