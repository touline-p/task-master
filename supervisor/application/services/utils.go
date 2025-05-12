package services

import (
	"fmt"
	"strings"
)

func ConcatenateErrors(errors []error) error {
	if len(errors) == 0 {
		return nil
	}

	if len(errors) == 1 {
		return errors[0]
	}

	errorMessages := make([]string, len(errors))
	for i, err := range errors {
		errorMessages[i] = err.Error()
	}

	return fmt.Errorf("multiple errors occurred: %s", strings.Join(errorMessages, "; "))
}
