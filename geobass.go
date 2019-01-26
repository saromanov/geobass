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

// New provides initialization of the app
func New() *GeoBass {
	return &GeoBass{
		items: make(map[uint64]interface{}),
	}
}

// Set provides setting to the cache
func (c *GeoBass) Set(p Point, value interface{}, expiration time.Duration) error {
	c.m.Lock()
	defer c.m.Unlock()
	hash, err := getHash(p)
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
	hash, err := getHash(p)
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
func getHash(p Point) (uint64, error) {
	hash, err := hashstructure.Hash(p.truncate(10), nil)
	if err != nil {
		return 0, fmt.Errorf("unable to hash point: %v", err)
	}
	return hash, nil
}
