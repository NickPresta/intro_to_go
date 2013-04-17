package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// START OMIT
const data = `[
	{"City": "Toronto", "Province": "Ontario", "Country": "Canada"},
	{"City": "Vancouver", "Province": "British Columbia", "Country": "Canada"}
]`

type GeoData struct {
	City     string
	Province string
	Country  string
}

func main() {
	var items []GeoData
	json.NewDecoder(strings.NewReader(data)).Decode(&items)
	for _, item := range items {
		fmt.Printf("%v, %v, %v\n",
			item.City, item.Province, item.Country)
	}
}

// STOP OMIT
