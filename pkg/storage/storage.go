package storage

import "errors"

// General errors for storage
var (
	ErrURLNotFound = errors.New("URL not found")
	ErrUrlExists   = errors.New("URL already exists")
)
