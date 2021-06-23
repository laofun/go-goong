# go-goong
 goong.io API wrappers for Golang


### Modules
- [X] Places
- [ ] Geocoding
- [ ] Directions

**Important:** This key should be kept secret on your server.

## Installation

To install the Go Client for Goong API, please execute the following `go get` command.

```bash
    go get github.com/laofun/go-goong
```

## Developer Documentation

View the [reference documentation](https://docs.goong.io/rest) for for API information.

## Requirements

- Go 1.14 or later.
- A Goong API key.

## Usage

```go
package main

import (
	"log"

	"github.com/kr/pretty"
	"github.com/laofun/go-goong"
	"github.com/laofun/go-goong/lib/places"
)

func main() {
	client, err := goong.NewClient("Insert-API-Key-Here")
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	r := &places.AutoCompleteOpts{
		Input: "quan 1",
	}
	resp, err := client.Places.Autocomplete(r)

	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	pretty.Println(resp)
}
```