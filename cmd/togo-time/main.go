package main

import (
	"fmt"
	"togotime"
)

func main() {
	h := togotime.NewAPIClient("API_TOKEN")
	fmt.Printf("%+v\n", h)
}
