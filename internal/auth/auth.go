package auth

import (
	"errors"
	"net/http"
	"strings"
)

// extracts api key from header of http request
// exmaple:
// Auth: ApiKey {insert apiKye here}
func GetApiKey(headers http.Header) (string, error) {

	val := headers.Get("Auth")

	if val == "" {
		return "", errors.New("no authentication info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed auth header")
	}

	return vals[1], nil

}
