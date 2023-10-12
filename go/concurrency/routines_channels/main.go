package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

const apiKey = "62037415837c13af5a1ceb8367e3514c"

func fetchWeather(city string, ch chan<- string, wg *sync.WaitGroup) interface{} {
	var data struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}

	defer wg.Done()

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&units=metric&appid=%s", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching weather for %s: %s", city, err)
		return data
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Printf("error decoding weather data for %s : %s", city, err)
		return data
	}
	ch <- fmt.Sprintf("%s temp is %v\n", city, data.Main.Temp)
	return data
}
func main() {
	start := time.Now()

	cities := []string{"Casablanca", "Rabat", "Marrakech", "Tanger"}

	ch := make(chan string)
	wg := sync.WaitGroup{}

	for _, city := range cities {
		wg.Add(1)
		go fetchWeather(city, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Println(result)
	}

	fmt.Println("The operation took :", time.Since(start))
}
