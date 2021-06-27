package services

import (
	"encoding/json"
	"flights-server/models"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func checkError(err error) {
	if err != nil {
		log.Print(err)
	}
}

//Fetch All Flight details from Sample Flight Database (JSON file)
func FetchAllFlightsService() (f *models.AllFlights) {

	file, err := os.OpenFile("database/flightList.json", os.O_RDWR|os.O_APPEND, 0666)
	checkError(err)

	var sampleflight models.AllFlights
	b, err := ioutil.ReadAll(file)
	checkError(err)

	if len(b) != 0 {
		err := json.Unmarshal(b, &sampleflight.Flights)
		checkError(err)
		return &sampleflight
	}

	return &sampleflight

}

//Fetch user provided source, destination flights
func FetchSourceAndDestinationFlightsService(input models.UserInput) models.AllFlights {

	allFlightData := FetchAllFlightsService()

	flights := FindFlights(*allFlightData, input)

	return flights
}

func FindFlights(allFlightData models.AllFlights, f models.UserInput) models.AllFlights {

	var finalList, srcOrDestMatchList models.AllFlights

	for i := 0; i < len(allFlightData.Flights); i++ {

		//Add flight details in array if there's no hope in between (direct flight)
		if strings.EqualFold(f.Origin, allFlightData.Flights[i].Origin) && strings.EqualFold(f.Destination, allFlightData.Flights[i].Destination) {

			layout := "15:04"
			departure := allFlightData.Flights[i].DepartureTime
			arrival := allFlightData.Flights[i].ArrivalTime

			t1, err := time.Parse(layout, departure)
			checkError(err)
			t2, err := time.Parse(layout, arrival)
			checkError(err)
			diff := t2.Sub(t1)

			data := models.Flight{
				Origin:        f.Origin,
				Destination:   f.Destination,
				Name:          allFlightData.Flights[i].Name,
				FlightNo:      allFlightData.Flights[i].FlightNo,
				ArrivalTime:   allFlightData.Flights[i].ArrivalTime,
				DepartureTime: allFlightData.Flights[i].DepartureTime,
				Date:          allFlightData.Flights[i].Date,
				TotalDuration: diff.String(),
			}

			finalList.Flights = append(finalList.Flights, data)
			continue
		} else {

			// find those flights whose origin / destination matches with user origin/ destination and added them in another array
			if strings.EqualFold(f.Origin, allFlightData.Flights[i].Origin) || strings.EqualFold(allFlightData.Flights[i].Destination, f.Destination) {
				srcOrDestMatchList.Flights = append(srcOrDestMatchList.Flights, allFlightData.Flights[i])
			}
		}

	}

	// Find out matching origin and destination from 2 array (hope's in between origin and destination)
	for i := 0; i < len(srcOrDestMatchList.Flights); i++ {
		for j := i + 1; j < len(srcOrDestMatchList.Flights); j++ {

			if strings.EqualFold(f.Origin, srcOrDestMatchList.Flights[i].Origin) && strings.EqualFold(srcOrDestMatchList.Flights[i].Destination, srcOrDestMatchList.Flights[j].Origin) && strings.EqualFold(srcOrDestMatchList.Flights[j].Destination, f.Destination) {

				var layover models.AllFlights

				layover.Flights = append(layover.Flights, srcOrDestMatchList.Flights[i])
				layover.Flights = append(layover.Flights, srcOrDestMatchList.Flights[j])

				data := models.Flight{
					Origin:        f.Origin,
					Destination:   f.Destination,
					Name:          allFlightData.Flights[i].Name,
					FlightNo:      allFlightData.Flights[i].FlightNo,
					ArrivalTime:   allFlightData.Flights[i].ArrivalTime,
					DepartureTime: allFlightData.Flights[i].DepartureTime,
					Layover:       layover.Flights,
				}

				finalList.Flights = append(finalList.Flights, data)
			}
		}

	}

	return finalList

}
