// Example of communicating with a goroutine function that loads the weather.
// Written by: Waqqas Sheikh (https://www.github.com/w-k-s)
// For: Dubai DevFest 2018 (https://www.meetup.com/en-AU/GDG-Dubai/events/253941428/)

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// You'll need to create an account at https://home.openweathermap.org/users/sign_up
// to get an API key and run the sample
const API_KEY = ""

func loadWeather(c chan string) {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=dubai&appid=" + API_KEY)
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	c <- string(bytes)
}

func countMilliseconds() {
	for counter := 1; ; counter++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(counter)
	}
}

func main() {
	weatherChannel := make(chan std ring)
	go loadWeather(weatherChannel)
	go countMilliseconds()
	fmt.Println(<-weatherChannel)
}
