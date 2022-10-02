package request

import (
	"fmt"
	"net/http"
)

func RequestForecast() *http.Response {
	site := "https://api.open-meteo.com/v1/forecast?latitude=59.3328&longitude=18.0645&hourly=temperature_2m"
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println(err)
	}

	return (resp)
}
