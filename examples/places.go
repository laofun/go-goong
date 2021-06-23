package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/kr/pretty"
	"github.com/laofun/go-goong"
	"github.com/laofun/go-goong/lib/places"
)

var (
	apiKey   = flag.String("api_key", "", "API Key for using Goong API.")
	input    = flag.String("input", "", "Your search keyword (Required)")
	location = flag.String("location", "", "Coordinates for location biased search. This must be specified as latitude,longitude.")
	radius   = flag.Int("radius", 0, "Limits Search to a radius from specified location (in km). Defaults to 50")
	limit    = flag.Int("limit", 0, "Limit number of results. Defaults to 10")
)

func main() {
	flag.Parse()
	var client *goong.Client
	var err error

	if *apiKey == "" {
		usageAndExit("Please specify an API Key, or Client ID and Signature.")
	}

	client, err = goong.NewClient(*apiKey)
	check(err)

	r := &places.AutoCompleteOpts{
		Input:    *input,
		Location: *location,
		Radius:   *radius,
		Limit:    *limit,
	}

	resp, err := client.Places.Autocomplete(context.Background(), r)
	check(err)

	pretty.Println("places.AutoComplete:", resp)

	if len(resp.Predictions) > 0 {
		detail, err := client.Places.Detail(context.Background(), &places.DetailOpts{
			PlaceID: resp.Predictions[0].PlaceID,
		})
		check(err)

		pretty.Println("places.Detail:", detail)
	}
}
func check(err error) {
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
}
func usageAndExit(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	fmt.Println("Flags:")
	flag.PrintDefaults()
	os.Exit(2)
}
