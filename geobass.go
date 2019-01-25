package geobass

import (
	"sync"
	"time"
)

// GeoBass defines app
type GeoBass struct {
	m     sync.RWMutex
	items map[Point]string
}

// Point defines point on store
type Point struct {
	Latitude  float64
	Longitude float64
}

// Set provides setting to the cache
func (c *GeoBass) Set(position Point, value interface{}, expiration time.Duration) {
	c.m.Lock()
	defer c.m.Unlock()
}

// Get provides getting object by the point
func (c *GeoBass) Get(p Point) (interface{}, error) {
	c.m.RLock()
	defer c.m.RUnlock()
	return nil, nil
}
