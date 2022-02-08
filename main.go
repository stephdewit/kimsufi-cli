package main

import (
	"flag"
	"fmt"
	"github.com/stephdewit/kimsufi-cli/availability"
	"os"
)

func main() {
	availabilityCommand := flag.NewFlagSet("availability", flag.ExitOnError)
	availabilityCountry := availabilityCommand.String("country", "fr", "Country")
	availabilityRegion := availabilityCommand.String("region", "europe", "Region")

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Missing sub-command\n")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "availability":
		availabilityCommand.Parse(os.Args[2:])

		if len(availabilityCommand.Args()) < 1 {
			fmt.Fprintf(os.Stderr, "Missing hardware parameter\n")
			os.Exit(1)
		}

		hardware := availabilityCommand.Args()[0]

		if hardware == "" {
			fmt.Fprintf(os.Stderr, "Empty hardware parameter\n")
		}

		availability, err := availability.Get(
			*availabilityCountry,
			*availabilityRegion,
			hardware)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}

		fmt.Println(availability)
	default:
		fmt.Fprintf(os.Stderr, "Expected 'availability' sub-command\n")
		os.Exit(1)
	}
}
