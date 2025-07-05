package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/CKroes97/go-gadgematron/internal/modules/timecheck"
	"github.com/CKroes97/go-gadgematron/internal/modules/weather"
)

var cityCoords = map[string]weather.Coord{
	"amsterdam": weather.Coord{
		Lat: 52.3676,
		Lon: 4.9041,
	},
	"rotterdam": weather.Coord{
		Lat: 51.9244,
		Lon: 4.4777,
	},
	"utrecht": weather.Coord{
		Lat: 52.0907,
		Lon: 5.1214,
	},
	"den haag": weather.Coord{
		Lat: 52.0705,
		Lon: 4.3007,
	},
	"groningen": weather.Coord{
		Lat: 53.2194,
		Lon: 6.5665,
	},
	"maastricht": weather.Coord{
		Lat: 50.8514,
		Lon: 5.6900,
	},
}

func ShowMainMenu() {
	fmt.Println("=== Go-Gadgematron ===")
	fmt.Println("Choose a module:")
	fmt.Println("1) Check time")
	fmt.Println("2) Weather report")
	fmt.Print("Enter choice: ")

	var choice int
	_, err := fmt.Scanln(&choice)
	if err != nil {
		fmt.Println("Invalid input:", err)
		return
	}

	switch choice {
	case 1:
		timecheck.Run()
	case 2:
		handleWeather()
	default:
		fmt.Println("Invalid choice")
	}
}

func handleWeather() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Dutch city name: ")
	cityInput, _ := reader.ReadString('\n')
	city := strings.ToLower(strings.TrimSpace(cityInput))

	coord, ok := cityCoords[city]
	if !ok {
		fmt.Println("Sorry, city not found in database.")
		return
	}

	caser := cases.Title(language.Dutch)
	cityTitle := caser.String(city)
	if err := weather.GetWeather(coord, cityTitle); err != nil {
		fmt.Println("Error:", err)
	}
}
