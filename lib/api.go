package lib

import (
	"encoding/json"
	"net/http"
	"os"
)

// get the API's response to a simple HTTP-GET request
func apiGet(url string, target interface{}) {
	r, err := http.Get(url)

	if err != nil {
		println("Unable to read from API:", err.Error())
		os.Exit(1)
	}

	err = json.NewDecoder(r.Body).Decode(&target)

	if err != nil {
		println("Unable to read from API:", err.Error())
		os.Exit(1)
	}

	r.Body.Close()
}
