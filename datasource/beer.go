package datasource

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"../model"
)

// GetBeers as objects, from file
func GetBeers() (beers []model.BeerDatasetElement) {
	var datasetElements []model.BeerDatasetElement

	jsonFile, err := os.Open("datasource/open-beer-database.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened open-beer-database.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(byteValue, &datasetElements)

	return datasetElements
}
	