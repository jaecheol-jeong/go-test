package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Weather is
type Weather struct {
	Coord struct {
		Lon float32 `json:"lon"`
		Lat float32 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		No          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Main struct {
		Temp      float32 `json:"temp"`
		Pressure  float32 `json:"pressure"`
		Humidity  int     `json:"humidity"`
		TempMin   float32 `json:"temp_min"`
		TempMax   float32 `json:"temp_max"`
		SeaLevel  float32 `json:"sea_level"`
		GrndLevel float32 `json:"srnd_level"`
	} `json:"main"`
	Wind struct {
		Speed float32 `json:"speed"`
		Deg   float32 `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  float32 `json:"dt"`
	Sys struct {
		Message float32 `json:"message"`
		Country string  `json:"country"`
		Sunrise float32 `json:"sunrise"`
		Sunset  float32 `json:"sunset"`
	} `json:"sys"`
	No   int    `json:"id"`
	Name string `json:"name"`
	Cod  int    `json:"cod"`
}

func main() {

	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=Seoul,KR&APPID=(APP_ID)&units=metric")

	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	println(string(data))

	var szData = []byte(data)
	var weather Weather
	err = json.Unmarshal(szData, &weather)
	if err != nil {
		panic(err)
	}

	fmt.Printf("최고 기온: %.2f °C\n", weather.Main.TempMax)
	fmt.Printf("최저 기온: %.2f °C\n", weather.Main.TempMin)
	fmt.Printf("현재기온: %.2f °C\n", weather.Main.Temp)

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}
