package main

import (
	"net/http"
	"strings"
	"testing"
)

func TestFlightAPI(t *testing.T) {

	input := strings.NewReader(`{
		Origin:      "Mumbai",
		Destination: "Delhi",
	}`)

	//Fetch All Flights
	t.Run("Test Fetch Flight API", func(t *testing.T) {

		res, err := http.Get("http://localhost:4700/o/flights")
		if err != nil {
			t.Errorf("unexpected error, got error:%s", err)
		}
		if res.StatusCode != 200 {
			t.Errorf("Expected status code 200, got status code %d:", res.StatusCode)
		}

	})

	//To check unauthorized access

	t.Run("Test source and destination flights", func(t *testing.T) {

		res, err := http.Post("http://localhost:4700/r/searchflights", "application/json", input)
		if err != nil {
			t.Errorf("unexpected error, got error:%s", err)
		}

		if res.StatusCode != 401 {
			t.Errorf("Expected status code 401, got status code %d:", res.StatusCode)
		}

	})

}
