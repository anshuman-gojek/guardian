package grant

import "errors"

var (
	ErrEmptyIDParam        = errors.New("grant id can't be empty")
	ErrGrantNotFound       = errors.New("grant not found")
	ErrEmptyImportedGrants = errors.New("imported grants not found")
)
