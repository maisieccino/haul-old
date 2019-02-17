package v1alpha1

import (
	"fmt"
)

// UnknownAPIVersionError is thrown when the
// wrong API version is parsed
type UnknownAPIVersionError struct {
	GivenVersion string
}

// UnknownAPIVersionErrorText is the format string
// for the error text raised when an unknown API version
// is parsed.
const UnknownAPIVersionErrorText = "Wrong API version. Expected: v1alpha1. Actual: %s"

func (err UnknownAPIVersionError) Error() string {
	return fmt.Sprintf(UnknownAPIVersionErrorText, err.GivenVersion)
}

// UnknownKindError is thrown when a kind is
// provided that this API does not know about
type UnknownKindError struct {
	GivenKind string
}

// UnknownKindErrorText is the format string for the error
// text raised when an unknown kind is provided
const UnknownKindErrorText = "Unknown Thing of kind %s"

func (err UnknownKindError) Error() string {
	return fmt.Sprintf(UnknownKindErrorText, err.GivenKind)
}

// FailedUnmarshalError is thrown when the program fails
// to parse the YAML thing manifest correctly.
type FailedUnmarshalError struct {
	string
}

// FailedUnmarshalErrorText is the formatted error text provided
// to the FailedUnmarshalError
const FailedUnmarshalErrorText = "Failed to parse YAML file: %s"

func (err FailedUnmarshalError) Error() string {
	return fmt.Sprintf(FailedUnmarshalErrorText, err.string)
}
