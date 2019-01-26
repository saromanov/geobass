package geobass

import (
	"fmt"
	"sync"
	"time"

	"github.com/mitchellh/hashstructure"
)

// GeoBass defines app
type GeoBass struct {
	m     sync.RWMutex
	items map[uint64]interface{}
}

// Point defines point on store
type Point struct {
	Latitude  float64
	Longitude float64
}

// Set provides setting to the cache
func (c *GeoBass) Set(p Point, value interface{}, expiration time.Duration) error {
	c.m.Lock()
	defer c.m.Unlock()
	hash, err := hashstructure.Hash(p, nil)
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
	return nil, nil
}
