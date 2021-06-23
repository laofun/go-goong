package goong

import (
	"github.com/laofun/go-goong/lib/base"
	"github.com/laofun/go-goong/lib/places"
)

// Client may be used to make requests to the Goong API
type Client struct {
	base   *base.Base
	Places *places.Places
}

// NewClient constructs a new Client Goong API
func NewClient(apiKey string) (*Client, error) {
	m := &Client{}
	// Create base instance
	base, err := base.NewBase(apiKey)
	if err != nil {
		return nil, err
	}
	m.base = base

	// Bind modules
	m.Places = places.NewPlaces(m.base)

	return m, nil
}
