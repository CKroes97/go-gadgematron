package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Coord struct {
	Lat float64
	Lon float64
}

type metResponse struct {
	Properties struct {
		Timeseries []struct {
			Time string `json:"time"`
			Data struct {
				Instant struct {
					Details struct {
						AirTemperature float64 `json:"air_temperature"`
						WindSpeed      float64 `json:"wind_speed"`
						WindDirection  float64 `json:"wind_from_direction"`
						Humidity       float64 `json:"relative_humidity"`
					} `json:"details"`
				} `json:"instant"`
			} `json:"data"`
		} `json:"timeseries"`
	} `json:"properties"`
}

func GetWeather(coord Coord, city string) error {
	url := fmt.Sprintf("https://api.met.no/weatherapi/locationforecast/2.0/compact?lat=%.4f&lon=%.4f", coord.Lat, coord.Lon)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("User-Agent", "go-gadgematron/1.0 - https://github.com/CKroes97/go-gadgematron")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("fetching weather: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("API error: %s", resp.Status)
	}

	var data metResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return fmt.Errorf("decoding response: %w", err)
	}

	if len(data.Properties.Timeseries) == 0 {
		return fmt.Errorf("no weather data available")
	}

	instant := data.Properties.Timeseries[0].Data.Instant.Details

	fmt.Printf("Weather in %s:\n", city)
	fmt.Printf("Temperature: %.1f °C\n", instant.AirTemperature)
	fmt.Printf("Wind: %.1f m/s from %.0f°\n", instant.WindSpeed, instant.WindDirection)
	fmt.Printf("Humidity: %.0f%%\n", instant.Humidity)

	return nil
}
