package domain

import "errors"

var (
	ErrExists      = errors.New("user exists")
	ErrNotFound    = errors.New("user not found")
	ErrBadPassword = errors.New("bad password")
)
