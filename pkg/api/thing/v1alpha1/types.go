package v1alpha1

import (
	"time"
)

// Thing repesents a thing object.
type Thing struct {
	APIVersion string        `yaml:"apiVersion"`
	Kind       string        `yaml:"kind"`
	Metadata   ThingMetadata `yaml:"metadata"`
}

// ThingMetadata represents metadata for a given thing
type ThingMetadata struct {
	DateCreated  time.Time `yaml:"dateCreated"`
	DateModified time.Time `yaml:"dateModified"`
	Owner        string    `yaml:"owner"`
	Labels       []string  `yaml:"labels"`
}

// RichText represents a thing of type richText.
// It describes a block of Markdown-formatted text.
type RichText struct {
	Thing `yaml:",inline"`
	Spec  RichTextSpec `yaml:"spec"`
}

// RichTextSpec defines the spec block of a richText
// thing.
type RichTextSpec struct {
	Content string `yaml:"content"`
}

// Image defines a thing of type image.
// It describes an image.
type Image struct {
	Thing `yaml:",inline"`
	Spec  ImageSpec `yaml:"spec"`
}

// ImageSpec defines the spec block of an image
// thing.
type ImageSpec struct {
	Path    string `yaml:"path"`
	AltText string `yaml:"altText"`
}
