package errors

import "fmt"

type ValidationError struct {
	Err error

	Reason string
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("Validation Error: %s", v.Reason)
}
