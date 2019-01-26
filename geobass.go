package geobass

import (
	"fmt"
	"sync"

	"github.com/mitchellh/hashstructure"
)

// Accuracy specifies accuracy of the distance
type Accuracy int32

const (
	// Range11KM The first decimal place is worth up to 11.1 km
	// eg: 30.1, 50.2
	Range11KM Accuracy = 1 + iota

	// Range1KM The second decimal place is worth up to 1.1 km
	// eg: 30.16, 50.23
	Range1KM

	// Range110M The third decimal place is worth up to 110 m
	// eg: 30.167, 50.231
	Range110M

	// Range11M The fourth decimal place is worth up to 11 m
	// eg: 30.1672, 50.2311
	Range11M

	// Range1M The fifth decimal place is worth up to 1.1 m
	// eg: 30.16722, 50.23118
	Range1M

	// Range11CM The sixth decimal place is worth up to 0.11 m
	// eg: 30.167221, 50.231189
	Range11CM

	// Range11MM The seventh decimal place is worth up to 11 mm
	// eg: 30.1672217, 50.2311896
	Range11MM

	// Range1MM The eighth decimal place is worth up to 1.1 mm
	// eg: 30.16722175, 50.23118962
	Range1MM
)

// GeoBass defines app
type GeoBass struct {
	m         sync.RWMutex
	items     map[uint64]interface{}
	precision Accuracy
}

// New provides initialization of the app
func New(precision Accuracy) *GeoBass {
	return &GeoBass{
		items:     make(map[uint64]interface{}),
		precision: precision,
	}
}

// Set provides setting to the cache
func (c *GeoBass) Set(p Point, value interface{}) error {
	c.m.Lock()
	defer c.m.Unlock()
	hash, err := getHash(p, c.precision)
	if err != nil {
		return fmt.Errorf("unable to hash point: %v", err)
	}
	c.items[hash] = value
	return nil
}

// Get provides getting object by the point
func (c *GeoBass) Get(p Point) (interface{}, error) {
	c.m.RLock()
	defer c.m.RUnlock()
	hash, err := getHash(p, c.precision)
	if err != nil {
		return nil, fmt.Errorf("unable to hash input point: %v", err)
	}
	value, ok := c.items[hash]
	if !ok {
		return nil, fmt.Errorf("unable to find element: %v", err)
	}
	return value, nil
}

// getHash retruns hash of the structure
func getHash(p Point, precision Accuracy) (uint64, error) {
	hash, err := hashstructure.Hash(p.truncate(precision), nil)
	if err != nil {
		return 0, fmt.Errorf("unable to hash point: %v", err)
	}
	return hash, nil
}
