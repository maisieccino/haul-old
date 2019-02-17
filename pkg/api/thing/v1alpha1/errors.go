package v1alpha1

import (
	"fmt"
)

// IncorrectAPIVersionError is thrown when the
// wrong API version is parsed.
type IncorrectAPIVersionError struct {
	GivenVersion string
}

func (err IncorrectAPIVersionError) Error() string {
	return fmt.Sprintf("Wrong API version. Expected: v1alpha1. Actual: %s", err.GivenVersion)
}
