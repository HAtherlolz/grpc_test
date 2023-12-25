package storage

import "errors"

var (
	ErrorUserExists   = errors.New("User already exists")
	ErrorUserNotFound = errors.New("User not found")
	ErrorAppNotFound  = errors.New("App not found")
)
