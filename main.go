package main

import (
	"fmt"
	"github.com/stephdewit/kimsufi-cli/availability"
	"os"
)

func main() {
	availability, err := availability.Get("fr", "europe", "2201sk011")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't get availability: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(availability)
}
