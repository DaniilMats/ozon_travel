package endpointHandlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DaniilMats/ozon_travel/ozonApi"
	"github.com/DaniilMats/ozon_travel/structs"
)

// FormHandler Функция FormHandler реализует интерфейс func(ResponseWriter, *Request),
// необходимый для обработки пост данных с формы form.html. Данные собираются
// для отправки graphql запроса и в конце формируют на фронте таблицу с тарифами.
func FormHandler(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		log.Fatal(writer, err)
	}
	date := request.FormValue("date")
	flightFrom := request.FormValue("flight_from")
	flightTo := request.FormValue("flight_to")

	airportFrom := ozonApi.GetAirportByCity(flightFrom)
	airportTo := ozonApi.GetAirportByCity(flightTo)
	flightData := ozonApi.GetFlightsByDatesAndAirports(date, airportFrom, airportTo)

	table := structs.TABLE_TEMPLATE
	for _, tariff := range flightData {
		template := `<tr><td>%s</td><td>%d</td><td>%d</td><td>%s</td><td>%s</td></tr>`
		for i := range tariff.Flights[0] {
			for j := range tariff.Flights[0][i].Segments {
				html := fmt.Sprintf(template, tariff.TariffDetails.Name, tariff.Price.Rub, tariff.Flights[0][i].AvailSeats, tariff.Flights[0][i].Segments[j].FlightNumber, tariff.Flights[0][i].Segments[j].Airline.Name)
				table += html
			}
		}
	}
	table += `</table></body></html>`
	_, err := fmt.Fprintln(writer, table)
	if err != nil {
		log.Fatal(err)
	}
}
