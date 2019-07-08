package model

import (
	"github.com/jinzhu/gorm"
)

// BeerDatasetElement ...
type BeerDatasetElement struct {
	gorm.Model
	DatasetID string `json:"datasetId"`
	RecordID string `json:"recordId"`
	Fields Beer `json:"fields"`
}
