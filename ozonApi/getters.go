package ozonApi

import (
	"awesomeProject/structs"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"strings"
)

// Функция getBytesResponse это обертка над стандартным пайплайном запроса
// с поправкой на то, что добавляется обязательный header и возвращается
// сразу слайс с байтами ответа.
func getBytesResponse(method, url string, payload *strings.Reader) []byte {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	return body
}

// Функция GetAirportByCity отправляет запрос на выгрузку всех аэропортов.
// По данным, полученным из ответа строится структура AirportSearch и отдается
// наружу код аэропорта. Всегда отдается первый аэропорт в списке, для облегчения
// бизнес-логике.
func GetAirportByCity(city string) string {
	url := "https://www.ozon.travel/graphql"
	method := "POST"
	queryTemplate := `{"query":"query FlightAirports($word: String!) {flightAirports(word: $word) {code airportName cityName}}",
						"variables":{"word":"%s"}}`
	query := fmt.Sprintf(queryTemplate, city)
	payload := strings.NewReader(query)

	responseBody := getBytesResponse(method, url, payload)
	var result structs.AirportSearch
	if err := json.Unmarshal(responseBody, &result); err != nil {
		fmt.Println(err)
	}
	aiportCode := result.Data.Airports[0].Code
	return aiportCode
}

// Функция GetFlightsByDatesAndAirports всех тарифов.
// Возвращает слайс структур с информацией о тарифах
func GetFlightsByDatesAndAirports(date, airportFrom, airportTo string) []structs.Tariff {
	url := "https://www.ozon.travel/graphql"
	method := "POST"
	queryTemplate := `{"query":"query FlightSearch($input: FlightSearchInput!) {flightSearch(input: $input) {  expires searchId tariffs {price {rub}  tariffDetails { name exchangeable refundable }  flights {duration availSeats  segments {flightNumber departureTime arrivalTime duration stopDuration airline { code name } availSeats }  }  }}}","variables":{"input":{"routes":[{"date":"%s","from":"%s","to":"%s"}],"adults":1,"children":0,"infants":0,"serviceClass":"ECONOMY","fromMetasearch":false,"noInstantSearch":false}}}`

	query := fmt.Sprintf(queryTemplate, date, airportFrom, airportTo)
	payload := strings.NewReader(query)

	responseBody := getBytesResponse(method, url, payload)
	rawTariffs := gjson.GetBytes(responseBody, "data.flightSearch.tariffs").Array()
	tariffs := getTariffs(rawTariffs)
	return tariffs
}

// Функция getTariffs формирует структуры Tariff и добавляет их в массив
func getTariffs(rawTariffs []gjson.Result) []structs.Tariff {
	var tariffs []structs.Tariff
	for _, tariff := range rawTariffs {
		strTariff := tariff.String()
		var result structs.Tariff
		if err := json.Unmarshal([]byte(strTariff), &result); err != nil {
			fmt.Println(err)
		}
		tariffs = append(tariffs, result)
	}
	return tariffs
}
