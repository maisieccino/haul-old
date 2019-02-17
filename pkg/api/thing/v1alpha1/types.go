package v1alpha1

import (
	"time"
)

// Thing represents a thing object
type Thing interface {
	GetAPIVersion() string
	GetKind() string
	GetMetadata() ThingMetadata
}

// ThingStruct represents a thing object.
type ThingStruct struct {
	APIVersion string        `yaml:"apiVersion"`
	Kind       string        `yaml:"kind"`
	Metadata   ThingMetadata `yaml:"metadata"`
}

// GetAPIVersion returns the API version of the thing.
func (th ThingStruct) GetAPIVersion() string {
	return th.APIVersion
}

// GetKind returns the kind (type) of the thing.
func (th ThingStruct) GetKind() string {
	return th.Kind
}

// GetMetadata returns the metadata of a struct.
func (th ThingStruct) GetMetadata() ThingMetadata {
	return th.Metadata
}

// ThingMetadata represents metadata for a given thing
type ThingMetadata struct {
	DateCreated  time.Time `yaml:"dateCreated"`
	DateModified time.Time `yaml:"dateModified"`
	Owner        string    `yaml:"owner"`
	Labels       []string  `yaml:"labels"`
}

// RichText defines a Rich Text object
type RichText interface {
	GetSpec() RichTextSpec
	Thing
}

// RichTextStruct represents a thing of type richText.
// It describes a block of Markdown-formatted text.
type RichTextStruct struct {
	ThingStruct `yaml:",inline"`
	Spec        RichTextSpec `yaml:"spec"`
}

// RichTextSpec defines the spec block of a richText
// thing.
type RichTextSpec struct {
	Content string `yaml:"content"`
}

// GetAPIVersion returns the API version of the thing.
func (rtc RichTextStruct) GetAPIVersion() string {
	return rtc.APIVersion
}

// GetKind returns the kind (type) of the thing.
func (rtc RichTextStruct) GetKind() string {
	return rtc.Kind
}

// GetMetadata returns the metadata of a struct.
func (rtc RichTextStruct) GetMetadata() ThingMetadata {
	return rtc.Metadata
}

// GetSpec returns the spec for the RichText thing
func (rtc RichTextStruct) GetSpec() RichTextSpec {
	return rtc.Spec
}

// Image defines a thing of kind image
type Image interface {
	GetSpec() ImageSpec
	Thing
}

// ImageStruct defines a thing of type image.
// It describes an image.
type ImageStruct struct {
	ThingStruct `yaml:",inline"`
	Spec        ImageSpec `yaml:"spec"`
}

// ImageSpec defines the spec block of an image
// thing.
type ImageSpec struct {
	Path    string `yaml:"path"`
	AltText string `yaml:"altText"`
}

// GetAPIVersion returns the API version of the thing.
func (i ImageStruct) GetAPIVersion() string {
	return i.APIVersion
}

// GetKind returns the kind (type) of the thing.
func (i ImageStruct) GetKind() string {
	return i.Kind
}

// GetMetadata returns the metadata of a struct.
func (i ImageStruct) GetMetadata() ThingMetadata {
	return i.Metadata
}

// GetSpec returns the spec of the image thing
func (i ImageStruct) GetSpec() ImageSpec {
	return i.Spec
}
