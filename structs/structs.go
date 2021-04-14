package structs

type Airport struct {
	Code string `json:"code"`
	Name string `json:"airportName"`
	City string `json:"cityName"`
}

type Airports struct {
	Airports []Airport `json:"flightAirports"`
}

type AirportSearch struct {
	Data Airports `json:"data"`
}

// структуры для запроса FlightSearch

type Airline struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Segment struct {
	FlightNumber  string  `json:"flightNumber"`
	DepartureTime string  `json:"departureTime"`
	ArrivalTime   string  `json:"arrivalTime"`
	Duration      int     `json:"duration"`
	StopDuration  int     `json:"stopDuration"`
	Airline       Airline `json:"airline"`
	AvailSeats    int     `json:"availSeats"`
}

type Flight struct {
	Duration   int       `json:"duration"`
	AvailSeats int       `json:"availSeats"`
	Segments   []Segment `json:"segments"`
}

type TariffDetails struct {
	Name         string `json:"name"`
	Exchangeable bool   `json:"exchangeable"`
	Refundable   bool   `json:"refundable"`
}

type Price struct {
	Rub int `json:"rub"`
}

type Tariff struct {
	Price         Price         `json:"price"`
	TariffDetails TariffDetails `json:"tariffDetails"`
	Flights       [][]Flight    `json:"flights"`
}
