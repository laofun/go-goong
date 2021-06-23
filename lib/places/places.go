// https://docs.goong.io/rest/place/

package places

import (
	"context"

	"github.com/google/go-querystring/query"
	"github.com/laofun/go-goong/lib/base"
)

const (
	apiName         = "place"
	apiAutocomplete = "autocomplete"
	apiDetail       = "detail"
)

// Places api wrapper instance
type Places struct {
	base *base.Base
}

// NewPlaces Create a new NewPlaces API wrapper
func NewPlaces(base *base.Base) *Places {
	return &Places{base}
}

type AutoCompleteOpts struct {
	Input    string `url:"input,omitempty"`
	Location string `url:"location,omitempty"`
	Radius   int    `url:"radius,omitempty"`
	Limit    int    `url:"limit,omitempty"`
}

type AutoCompleteResponse struct {
	Predictions []base.Predictions `json:"predictions"`
	Status      string             `json:"status"`
}

// Places Search by keyword with autocomplete
// Return predictions according to search keyword
func (g *Places) Autocomplete(ctx context.Context, req *AutoCompleteOpts) (*AutoCompleteResponse, error) {
	v, err := query.Values(req)
	if err != nil {
		return nil, err
	}
	resp := AutoCompleteResponse{}
	err = g.base.Query(ctx, apiName, apiAutocomplete, &v, &resp)

	return &resp, err
}

type DetailOpts struct {
	PlaceID string `url:"place_id,omitempty"`
}

type DetailResponse struct {
	Result base.PlaceDetailResult `json:"result"`
	Status string                 `json:"status"`
}

// Get place detail by Id
// Return detail of a place by it's place_id
func (g *Places) Detail(ctx context.Context, req *DetailOpts) (*DetailResponse, error) {
	v, err := query.Values(req)
	if err != nil {
		return nil, err
	}
	resp := DetailResponse{}

	err = g.base.Query(ctx, apiName, apiDetail, &v, &resp)

	return &resp, err
}
