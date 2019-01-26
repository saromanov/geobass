package geobass

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeoBass(t *testing.T) {
	g := New(Range11KM)
	err := g.Set(Point{
		Latitude:  30.1672,
		Longitude: 50.2311,
	}, "data")
	assert.NoError(t, err)
	value, err := g.Get(Point{
		Latitude:  30.1670,
		Longitude: 50.2311,
	})
	assert.NoError(t, err)
	assert.Equal(t, value, "data")
	g.Clear()
}

func TestGeoBass2(t *testing.T) {
	g := New(Range110M)
	err := g.Set(Point{
		Latitude:  30.1672,
		Longitude: 50.2311,
	}, "data")
	assert.NoError(t, err)
	_, err = g.Get(Point{
		Latitude:  50.1670,
		Longitude: 50.2311,
	})
	assert.Error(t, err)
	g.Clear()
}
