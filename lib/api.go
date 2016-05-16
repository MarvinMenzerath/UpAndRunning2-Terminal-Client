package lib

import (
	"encoding/json"
	"net/http"
	"os"
)

// get the API's response to a simple HTTP-GET request
func apiGet(url string, target interface{}) error {
	r, err := http.Get(url)
	defer r.Body.Close()

	if err != nil {
		println("Unable to read from API.")
		os.Exit(1)
	}

	return json.NewDecoder(r.Body).Decode(&target)
}
