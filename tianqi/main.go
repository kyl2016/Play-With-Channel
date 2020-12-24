package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// https://siongui.github.io/2018/10/27/auto-detect-and-convert-html-encoding-to-utf8-in-go/

const file = "weather.json"

func main() {
	m := map[string]map[string]Weather{}
	data, err := ioutil.ReadFile(file)
	if err == nil {
		json.Unmarshal(data, &m)
	}

	for {
		resp, err := http.Get("https://tianqiapi.com/api?version=v6&appid=74959191&appsecret=9iFgGUMd")
		if err != nil {
			panic(err)
		}
		data, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		// b := transcode.FromByteArray(data).Decode("GBK").ToByteArray()

		var today Weather
		err = json.Unmarshal(data, &today)
		if err != nil {
			panic(err)
		}

		if _, ok := m[today.CityID]; !ok {
			m[today.CityID] = map[string]Weather{}
		}

		yesterday := getYesterdayWeather(m, today)

		rule := Rule{TemperatureDiff: 1, MinPM25: 100, WinSpeed: 3}
		if satisfy(yesterday, today, rule) {
			alarm := createAlarm(yesterday, today)
			send(alarm)
		}

		m[today.CityID][today.Date] = today
		result, _ := json.MarshalIndent(m, "", "\t")
		ioutil.WriteFile(file, result, os.ModePerm)

		time.Sleep(time.Hour * 24)
	}
}

func createAlarm(yesterday *Weather, today Weather) string {
	minTempRise := "⬆"
	minTempDiff := int(math.Abs(float64(parseInt(yesterday.MinTemperature) - parseInt(today.MinTemperature))))
	if parseInt(yesterday.MinTemperature) > parseInt(today.MinTemperature) {
		minTempRise = "⬇"
	}
	maxTempRise := "⬆"
	maxTempDiff := int(math.Abs(float64(parseInt(yesterday.MaxTemperature) - parseInt(today.MaxTemperature))))
	if parseInt(yesterday.MaxTemperature) > parseInt(today.MaxTemperature) {
		maxTempRise = "⬇"
	}

	// 最低气温 %s°，比昨天%s %2d°\n最高气温 %s°，比昨天%s %2d°\n
	return fmt.Sprintf("%s\n气温：%s°%s%d° ~ %s°%s%d°\n%s %s \nPM2.5 %s",
		// yesterday.MinTemperature, yesterday.MaxTemperature,
		today.Date,
		today.MinTemperature, minTempRise, minTempDiff, today.MaxTemperature, maxTempRise, maxTempDiff, today.Win, today.WinSpeed, today.PM25)
}

func send(alarm string) {
	http.Post("https://hooks.slack.com/services/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", "application/json", strings.NewReader(fmt.Sprintf(`{"text":"%s @channel"}`, alarm)))
	fmt.Println(alarm)
}

type Rule struct {
	TemperatureDiff int
	MinPM25         int
	WinSpeed        int
}

func parseInt(v string) int {
	r, err := strconv.Atoi(v)
	panicIfNotNil(err)
	return r
}

func panicIfNotNil(err error) {
	if err != nil {
		panic(err)
	}
}

func getYesterdayWeather(m map[string]map[string]Weather, today Weather) *Weather {
	d, _ := time.Parse("2006-01-02", today.Date)
	yesterday := d.Add(-time.Hour * 24).Format("2006-01-02")

	if city, ok := m[today.CityID]; ok {
		if yesterdayWeather, ok := city[yesterday]; ok {
			return &yesterdayWeather
		}
	}

	return nil
}

func satisfy(yesterday *Weather, today Weather, rule Rule) bool {
	if parseInt(strings.ReplaceAll(today.WinSpeed, "级", "")) >= rule.WinSpeed || parseInt(today.PM25) >= rule.MinPM25 {
		return true
	}

	if yesterday != nil {
		y, _ := strconv.Atoi(yesterday.MinTemperature)
		t, _ := strconv.Atoi(today.MinTemperature)

		if math.Abs(float64(y-t)) > float64(rule.TemperatureDiff) {
			return true
		}
	}
	return false
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
