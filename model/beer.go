package model

// BeerDatasetElement ...
type BeerDatasetElement struct {
	DatasetID string `json:"datasetId"`
	RecordID string `json:"recordId"`
	Fields Beer `json:"fields"`
}

// Beer ...
type Beer struct {
	ID string `json:"id"`
	Name string `json:"name"`
	ABV float32 `json:"abv"`
	BreweryID string `json:"brewery_id"`
	BreweryName string `json:"name_breweries"`
	City string `json:"city"`
	Country string `json:"country"`
	State string `json:"state"`
}

