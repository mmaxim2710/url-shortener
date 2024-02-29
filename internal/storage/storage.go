package storage

import "errors"

var (
	ErrURLNotFound      = errors.New("requested URL not found")
	ErrURLAlreadyExists = errors.New("requested URL already exists")
)
