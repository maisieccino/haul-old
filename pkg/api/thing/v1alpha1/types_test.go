package v1alpha1_test

import (
	"testing"
	"time"

	thing "github.com/mbellgb/haul/pkg/api/thing/v1alpha1"
	"github.com/stretchr/testify/assert"
)

var richTextYaml = `apiVersion: thing.haul.io/v1alpha1
kind: richText
metadata:
  dateCreated: 2018-12-14T20:45:00Z
  dateModified: 2018-12-14T20:45:00Z
  owner: 'Matt'
  bag: bag/name
  labels:
  - label-name
  - other-label
spec:
  content: hello this is a test`

var imgYaml = `apiVersion: thing.haul.io/v1alpha1
kind: image
metadata:
  dateCreated: 2018-12-14T20:45:00Z
  dateModified: 2018-12-14T20:45:00Z
  owner: 'Matt'
  bag: bag/name
  labels:
  - label-name
  - other-label
spec:
  path: /storage/images/hello.png
  altText: This is a photo.`

func TestThingUnmarshal(t *testing.T) {
	th, err := thing.UnmarshalThing([]byte(richTextYaml))
	assert.NoError(t, err)

	assert.Equal(t, th.GetAPIVersion(), "thing.haul.io/v1alpha1")
	assert.Equal(t, th.GetKind(), "richText")
	assert.Equal(t, th.GetMetadata().DateCreated.Month(), time.December)
	assert.Equal(t, th.GetMetadata().DateCreated.Hour(), 20)
	assert.Equal(t, len(th.GetMetadata().Labels), 2)
	assert.Equal(t, th.GetMetadata().Labels[0], "label-name")
}

func TestRichTextUnmarshal(t *testing.T) {
	th, err := thing.UnmarshalThing([]byte(richTextYaml))
	assert.NoError(t, err)
	text, ok := th.(thing.RichText)
	assert.True(t, ok)

	assert.Equal(t, th.GetAPIVersion(), "thing.haul.io/v1alpha1")
	assert.Equal(t, text.GetSpec().Content, "hello this is a test")
}

func TestImageUnmarshal(t *testing.T) {
	th, err := thing.UnmarshalThing([]byte(imgYaml))
	assert.NoError(t, err)
	img, ok := th.(thing.Image)
	assert.True(t, ok)

	assert.Equal(t, th.GetAPIVersion(), "thing.haul.io/v1alpha1")
	assert.Equal(t, img.GetSpec().Path, "/storage/images/hello.png")
	assert.Equal(t, img.GetSpec().AltText, "This is a photo.")
}

func TestCantUnmarshalWrongAPI(t *testing.T) {
	wrongAPIVersion := []byte(`apiVersion: thing.haul.io/v2
kind: richText
metadata: {}`)
	_, err := thing.UnmarshalThing(wrongAPIVersion)
	assert.Error(t, thing.IncorrectAPIVersionError{}, err)
}
