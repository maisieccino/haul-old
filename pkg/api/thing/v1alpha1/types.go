package v1alpha1

import (
	"time"
)

// Thing repesents an unmarshalled Thing file
type Thing struct {
	APIVersion string        `yaml:"apiVersion"`
	Kind       string        `yaml:"kind"`
	Metadata   ThingMetadata `yaml:"metadata"`
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

// ThingMetadata represents metadata for a given thing
type ThingMetadata struct {
	DateCreated  time.Time `yaml:"dateCreated"`
	DateModified time.Time `yaml:"dateModified"`
	Owner        string    `yaml:"owner"`
	Labels       []string  `yaml:"labels"`
}
