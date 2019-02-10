package v1alpha1_test

import (
	thing "github.com/mbellgb/haul/pkg/api/thing/v1alpha1"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"testing"
	"time"
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

func TestThingUnmarshal(t *testing.T) {
	th := thing.Thing{}
	err := yaml.Unmarshal([]byte(richTextYaml), &th)
	assert.NoError(t, err)

	assert.Equal(t, th.APIVersion, "thing.haul.io/v1alpha1")
	assert.Equal(t, th.Kind, "richText")
	assert.Equal(t, th.Metadata.DateCreated.Month(), time.December)
	assert.Equal(t, th.Metadata.DateCreated.Hour(), 20)
	assert.Equal(t, len(th.Metadata.Labels), 2)
	assert.Equal(t, th.Metadata.Labels[0], "label-name")
}

func TestRichTextUnmarshal(t *testing.T) {
	text := thing.RichText{}
	th := thing.Thing{}
	err := yaml.Unmarshal([]byte(richTextYaml), &text)
	assert.NoError(t, err)
	err = yaml.Unmarshal([]byte(richTextYaml), &th)
	assert.NoError(t, err)

	assert.Equal(t, text.APIVersion, "thing.haul.io/v1alpha1")
	assert.Equal(t, text.Spec.Content, "hello this is a test")
}
