package models

type Flight struct {
	FlightNo      string   `json:"flightNo"`
	Name          string   `json:"name"`
	Origin        string   `json:"origin"`
	Destination   string   `json:"destination"`
	Date          string   `json:"date"`
	Price         int      `json:"price"`
	DepartureTime string   `json:"depTime"`
	ArrivalTime   string   `json:"arrTime"`
	TotalDuration string   `json:"totalDuration"`
	Layover       []Flight `json:"layover"`
}

type AllFlights struct {
	Flights []Flight
}

type UserInput struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
}

type Result struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
