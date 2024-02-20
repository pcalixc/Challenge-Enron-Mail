package errors

import (
	"errors"
	"fmt"
)

type NotFoundError struct {
	File string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("file %q not found", e.File)
}

func Open(file string) error {
	return &NotFoundError{File: file}
}

// package bar

// if err := foo.Open("testfile.txt"); err != nil {
// 	var notFound *NotFoundError
// 	if errors.As(err, &notFound) {
// 	  // handle the error
// 	} else {
// 	  panic("unknown error")
// 	}
//   }

var (
	// The following two errors are exported
	// so that users of this package can match them
	// with errors.Is.

	ErrBrokenLink   = errors.New("link is broken")
	ErrCouldNotOpen = errors.New("could not open")

	// This error is not exported because
	// we don't want to make it part of our public API.
	// We may still use it inside the package
	// with errors.Is.

	errNotFound = errors.New("not found")
)
