package features

type HourlyResponse struct {
	Response *Response `json:"response"`
	Hourly   []Hourly  `json:"hourly_forecast"`
}

type Hourly struct {
	FCTTime       FCTTime   `json:"FCTTIME"`
	Condition     string    `json:"condition"`
	Icon          string    `json:"icon"`
	IconUrl       string    `json:"icon_url"`
	Temperature   Unit      `json:"temp"`
	WindSpeed     Unit      `json:"wspd"`
	WindDirection Direction `json:"wdir`
	FeelsLike     Unit      `json:"feelslike`
}

type Unit struct {
	English string `json:"english"`
	Metric  string `json:"metric"`
}

type Direction struct {
	Dir     string `json:"dir"`
	Degrees string `json:"degrees"`
}

type FCTTime struct {
	Epoch    string `json:"epoch"`
	DateTime string `json:"pretty"`
	Civil    string `json:"civil"`
}
