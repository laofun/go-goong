package base

import (
	"errors"
)

// ErrorAPIUnauthorized indicates authorization failed
var ErrorAPIUnauthorized = errors.New("Goong API error unauthorized")

// ErrorAPILimitExceeded indicates the API limit has been exceeded
var ErrorAPILimitExceeded = errors.New("Goong API error api rate limit exceeded")
