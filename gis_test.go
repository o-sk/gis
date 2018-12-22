package gis

import "testing"
import "github.com/stretchr/testify/assert"

func TestSearch(t *testing.T) {
	images, err := Search("猫")
	assert.NotNil(t, images)
	assert.Nil(t, err)

	image := images[0]
	assert.NotNil(t, image.ID)
	assert.NotNil(t, image.Destination)
	assert.NotNil(t, image.Source)
	assert.NotNil(t, image.Thumbnail)
}
