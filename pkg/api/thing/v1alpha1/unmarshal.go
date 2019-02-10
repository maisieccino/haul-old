package v1alpha1

import (
	"gopkg.in/yaml.v2"
)

// Unmarshal takes a YAML definition of a thing
// and returns the Go struct representation of it.
func Unmarshal(input []byte) (th Thing, err error) {
	err = yaml.Unmarshal(input, &th)
	return
}
