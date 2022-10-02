package main

import (
	"fmt"

	"github.com/noalerch/weatherman/request"
)

func main() {
	fmt.Println("Hello, world!")

	resp := request.RequestForecast()
	fmt.Println(resp)
}
