package utils

import (
	"errors"
)

// Use errors.Join when upgrading to Go 1.2	0
func Join(errs ...error) error {
	var errStr string
	for _, err := range errs {
		if err != nil {
			if len(errStr) > 0 {
				errStr += ", "
			}
			errStr += err.Error()
		}
	}
	if len(errStr) > 0 {
		return errors.New(errStr)
	}
	return nil
}
