package goong

import (
	"github.com/laofun/go-goong/lib/base"
	"github.com/laofun/go-goong/lib/places"
)

// Goong API Wrapper structure
type Goong struct {
	base   *base.Base
	Places *places.Places
}

// NewGoong Create a new Goong API instance
func NewGoong(apiKey string) (*Goong, error) {
	m := &Goong{}
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
