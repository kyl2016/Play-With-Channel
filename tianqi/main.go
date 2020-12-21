package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

// https://siongui.github.io/2018/10/27/auto-detect-and-convert-html-encoding-to-utf8-in-go/

func main() {
	for {

		resp, err := http.Get("https://tianqiapi.com/api?version=v6&appid=74959191&appsecret=9iFgGUMd")
		if err != nil {
			panic(err)
		}
		data, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		// b := transcode.FromByteArray(data).Decode("GBK").ToByteArray()

		var w Weather
		err = json.Unmarshal(data, &w)
		if err != nil {
			panic(err)
		}

		m := map[string]map[string]Weather{}
		if _, ok := m[w.CityID]; !ok {
			m[w.CityID] = map[string]Weather{}
		}

		m[w.CityID][w.Date] = w

		d, _ := time.Parse("2016-01-02", w.Date)
		yesterday := d.Add(-time.Hour * 24).Format("2016-01-02")

		if yesterdayWeather, ok := m[w.CityID][yesterday]; ok {
			y, _ := strconv.Atoi(yesterdayWeather.MinTemperature)
			t, _ := strconv.Atoi(w.MinTemperature)

			if y-t > 5 {
				// notify
				fmt.Println("温度骤降 ", (y - t), "°")
			}
		}

		result, _ := json.Marshal(m)
		ioutil.WriteFile("weather.json", result, os.ModePerm)

		time.Sleep(time.Hour * 24)
	}
}

type Weather struct {
	CityID             string `json:"cityid"`
	Date               string `json:"date"`
	CurrentTemperature string `json:"tem"`
	MaxTemperature     string `json:"tem1"`
	MinTemperature     string `json:"tem2"`
	UpdateTime         string `json:"update_time"`
	Desc               string `json:"wea"`
	Win                string `json:"win"`
	WinSpeed           string `json:"win_speed"`
	Humidity           string `json:"humidity"`
	PM25               string `json:"air_pm25"`
}
