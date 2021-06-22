package base

type StructuredFormatting struct {
	MainText      string `json:"main_text"`
	SecondaryText string `json:"secondary_text"`
}

type PlusCode struct {
	CompoundCode string `json:"compound_code"`
	GlobalCode   string `json:"global_code"`
}

type Predictions struct {
	Description          string               `json:"description"`
	MatchedSubstrings    []interface{}        `json:"matched_substrings"`
	PlaceID              string               `json:"place_id"`
	Reference            string               `json:"reference"`
	StructuredFormatting StructuredFormatting `json:"structured_formatting"`
	Terms                []interface{}        `json:"terms"`
	HasChildren          bool                 `json:"has_children"`
	DisplayType          string               `json:"display_type"`
	Score                float64              `json:"score"`
	PlusCode             PlusCode             `json:"plus_code"`
}

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
type Geometry struct {
	Location Location `json:"location"`
}
type PlaceDetailResult struct {
	PlaceID          string   `json:"place_id"`
	FormattedAddress string   `json:"formatted_address"`
	Geometry         Geometry `json:"geometry"`
	Name             string   `json:"name"`
}
