package availability

import (
	"fmt"
	"github.com/stephdewit/kimsufi-lib/availability"
)

func Get(country string, region string, hardware string) (string, error) {
	kimsufi := availability.Client{Country: country}
	availability, err := kimsufi.GetAvailability(region, hardware)
	if err != nil {
		return "", fmt.Errorf("Can't get availability: %w", err)
	}

	return availability, nil
}
