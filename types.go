package main

// 响应
type Response struct {
	Status   int    `json:"status"`
	CityName string `json:"city"`
	Data     Data   `json:"data"`
	Date     string `json:"date"`
	Message  string `json:"message"`
	Count    int    `json:"count"`
}

// 响应数据
type Data struct {
	ShiDu     string `json:"shidu"`
	Quality   string `json:"quality"`
	Ganmao    string `json:"ganmao"`
	Yesterday Day    `json:"yesterday"`
	Forecast  []Day  `json:"forecast"`
}

// 某一天的数据
type Day struct {
	Date    string  `json:"date"`
	Sunrise string  `json:"sunrise"`
	High    string  `json:"high"`
	Low     string  `json:"low"`
	Sunset  string  `json:"sunset"`
	Aqi     float32 `json:"aqi"`
	Fx      string  `json:"fx"`
	Fl      string  `json:"fl"`
	Type    string  `json:"type"`
	Notice  string  `json:"notice"`
}
