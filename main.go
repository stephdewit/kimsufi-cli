package main

import (
	"fmt"
	"github.com/stephdewit/kimsufi-lib/availability"
	"os"
)

func main() {
	kimsufi := availability.Client{Country: "fr"}
	availability, err := kimsufi.GetAvailability("europe", "2201sk091")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't get availability: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(availability)
}
