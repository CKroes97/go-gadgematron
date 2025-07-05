package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/CKroes97/go-gadgematron/internal/modules/timecheck"
	"github.com/CKroes97/go-gadgematron/internal/modules/weather"
)

var cityCoords = map[string]weather.Coord{
	"amsterdam":  {52.3676, 4.9041},
	"rotterdam":  {51.9244, 4.4777},
	"utrecht":    {52.0907, 5.1214},
	"den haag":   {52.0705, 4.3007},
	"groningen":  {53.2194, 6.5665},
	"maastricht": {50.8514, 5.6900},
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

	if err := weather.GetWeather(coord, strings.Title(city)); err != nil {
		fmt.Println("Error:", err)
	}
}
