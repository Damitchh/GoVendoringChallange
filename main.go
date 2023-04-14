package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type WeatherData struct {
	Water       int    `json:"water"`
	Wind        int    `json:"wind"`
	StatusWater string `json:"status_water"`
	StatusWind  string `json:"status_wind"`
}

func main() {

	for {
		water := getRandomValue()
		wind := getRandomValue()
		statusWater := getStatusWater(water)
		statusWind := getStatusWind(wind)

		data := WeatherData{
			Water:       water,
			Wind:        wind,
			StatusWater: statusWater,
			StatusWind:  statusWind,
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			return
		}

		resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("Error sending request:", err)
			return
		}

		if resp.StatusCode == http.StatusCreated {
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Error reading response body:", err)
				continue
			}
			fmt.Println("Response body:", string(body))
		} else {
			fmt.Println("Error: response status code", resp.StatusCode)
		}

		time.Sleep(15 * time.Second)
	}

}

func getRandomValue() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100) + 1
}

func getStatusWater(water int) string {
	switch {
	case water < 5:
		return "aman"
	case water < 9:
		return "siaga"
	default:
		return "bahaya"
	}
}

func getStatusWind(wind int) string {
	switch {
	case wind < 6:
		return "aman"
	case wind < 16:
		return "siaga"
	default:
		return "bahaya"
	}
}
