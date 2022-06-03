package errors

import "fmt"

// Custom Error interface
type MyError struct {
	Inner      error
	StatusCode int
	Message    string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("error caused due to %v; Status Code: %d: message: %v", e.Inner, e.StatusCode, e.Message)
}

func (e *MyError) Unwrap() error {
	return e.Inner
}
