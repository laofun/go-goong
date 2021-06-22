# go-goong
 goong.io API wrappers for Golang

See [here](https://docs.goong.io/rest) for for API information.

### Modules
- [X] Places
- [ ] Geocoding
- [ ] Directions

## Examples
### Initialisation
```go
import (
  "github.com/laofun/go-goong/lib"
  "github.com/laofun/go-goong/lib/base"
)
// Fetch apikey from somewhere
apiKey := os.Getenv("GOONG_API_KEY")

// Create new MAP instance
client := goong.NewGoong(apiKey)

```
### Places API
```go
import (
  "github.com/laofun/go-goong/lib/places"
)

placeData, err := client.Places.Autocomplete(&places.AutoCompleteOpts{
  Input: "quan 1"
})
```