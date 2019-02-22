package v1alpha1

import (
	"gopkg.in/yaml.v2"
)

// UnmarshalThing takes a YAML definition of a thing
// and returns the Go struct representation of it.
func UnmarshalThing(input []byte) (Thing, error) {
	// Parse a generic thing to obtain the Kind of the
	// Thing.
	var th ThingStruct
	if err := yaml.Unmarshal(input, &th); err != nil {
		return nil, err
	}

	// Check for correct API version
	if th.GetAPIVersion() != apiVersionString {
		return nil, UnknownAPIVersionError{GivenVersion: th.GetAPIVersion()}
	}

	switch th.GetKind() {
	case "richText":
		var rt RichTextStruct
		if err := yaml.Unmarshal(input, &rt); err != nil {
			return nil, err
		}
		return rt, nil
	case "image":
		var img ImageStruct
		if err := yaml.Unmarshal(input, &img); err != nil {
			return nil, err
		}
		return img, nil
	default:
		return nil, UnknownKindError{GivenKind: th.GetKind()}
	}

}
