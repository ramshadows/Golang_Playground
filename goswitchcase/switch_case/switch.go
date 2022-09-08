package goswitchcase

import (
	"fmt"
	"math/rand"
	"time"
)

func GoSwitchCase() {
	sec := time.Now().Unix()
	rand.Seed(sec)
	i := rand.Int31n(10)

	switch i {
	case 0:
		fmt.Print("zero...")
	case 1:
		fmt.Print("one...")
	case 2:
		fmt.Print("two...")

	default:
		fmt.Println("No match found")
	}

}

func Location(city string) (string, string) {
	var region string
	var continent string

	switch city {
	case "Kampala", "Darlesalam", "Nairobi":
		region, continent = "East Africa", "Africa"

	case "Kigali", "Burundi", "Bujumbura":
		region, continent = "Central Africa", "Africa"

	case "Irvine", "Los Angeles", "San Diego":
		region, continent = "California", "USA"
	default:
		region, continent = "Unknown", "Unknown"
	}

	return region, continent
}
