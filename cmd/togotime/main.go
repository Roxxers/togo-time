package main

import (
	"fmt"
	"togotime"
)

func main() {
	h := togotime.NewAPIClient("2136a22845735689d7a6e7f15e56e087")
	fmt.Printf("%+v\n", h.Workspaces[0].Projects)
}
