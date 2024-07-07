package gmap

import (
	"context"
	"log"

	"googlemaps.github.io/maps"
)

type GmapApiManager struct {
	client *maps.Client
}

type Location struct {
	Description string `json:"description"`

	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func NewGmapApiManager() *GmapApiManager {
	return &GmapApiManager{}
}

func (g *GmapApiManager) InitGmapApiManager(apiKey string) {
	client, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}

	g.client = client
}

func (g *GmapApiManager) GetLocationAutocomplete(input string) ([]Location, error) {
	r := &maps.PlaceAutocompleteRequest{
		Input: input,
	}

	places, err := g.client.PlaceAutocomplete(context.Background(), r)
	if err != nil {
		return nil, err
	}

	var locations []Location

	for _, value := range places.Predictions {
		var location Location

		location.Description = value.Description

		r := &maps.GeocodingRequest{
			Address: value.Description,
		}
		res, err := g.client.Geocode(context.Background(), r)
		if err != nil || len(res) == 0 {
			location.Latitude = 0
			location.Longitude = 0
		}

		location.Latitude = res[0].Geometry.Location.Lat
		location.Longitude = res[0].Geometry.Location.Lng

		locations = append(locations, location)
	}

	return locations, nil
}
