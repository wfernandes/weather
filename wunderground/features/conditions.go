package features

type ConditionsResponse struct {
	Response  Response   `json:"response"`
	Condition Conditions `json:"current_observation"`
}

type Conditions struct {
	Image                 Image               `json:"image"`
	DisplayLocation       DisplayLocation     `json:"display_location"`
	ObservationLocation   ObservationLocation `json:"observation_location"`
	StationId             string              `json:"station_id"`
	ObservationTimeRFC822 string              `json:"observation_time_rfc822"`
	ObservationEpoch      string              `json:"observation_epoch"`
	LocalTimeRFC822       string              `json:"local_time_rfc822"`
	LocalEpoch            string              `json:"local_epoch"`
	LocalTZShort          string              `json:"local_tz_short"`
	LocalTZOffset         string              `json:"local_tz_offset"`
	Weather               string              `json:"weather"`
	TempF                 float64             `json:"temp_f"`
	TempC                 float64             `json:"temp_c"`
	RelativeHumidity      string              `json:"relative_humidity"`
	Wind                  string              `json:"wind_string"`
	WindDir               string              `json:"wind_dir"`
	WindDegrees           float64             `json:"wind_degrees"`
	WindMPH               float64             `json:"wind_mph"`
	WindGustMPH           string             `json:"wind_gust_mph"`
	WindKPH               float64             `json:"wind_kph"`
	WindGustKPH           string             `json:"wind_gust_kph"`
	PressureMB            string              `json:"pressure_mb"`
	PressureIN            string              `json:"pressure_in"`
	PressureTrend         string              `json:"pressure_trend"`
	DewpointF             float64             `json:"dewpoint_f"`
	DewpointC             float64             `json:"dewpoint_c"`
	FeelslikeF            string              `json:"feelslike_f"`
	FeelslikeC            string              `json:"feelslike_c"`
	VisibilityMI          string              `json:"visibility_mi"`
	VisibilityKM          string              `json:"visibility_km"`
	Precip1hr             string              `json:"precip_1hr_string"`
	Precip1hrIN           string              `json:"precip_1hr_in"`
	Precip1hrMETRIC       string              `json:"precip_1hr_metric"`
	PrecipToday           string              `json:"precip_today_string"`
	PrecipTodayIN         string              `json:"precip_today_in"`
	PrecipTodayMETRIC     string              `json:"precip_today_metric"`
	Icon                  string              `json:"icon"`
	IconURL               string              `json:"icon_url"`
	ForecastURL           string              `json:"forecast_url"`
	HistoryURL            string              `json:"history_url"`
	ObURL                 string              `json:"ob_url"`
}

type DisplayLocation struct {
	City       string `json:"city"`
	State      string `json:"state"`
	StateName string `json:"state_name"`
	Country    string `json:"country"`
	Zip        string `json:"zip"`
	Latitude   string `json:"latitude"`
	Longitude  string `json:"longitude"`
	Elevation  string `json:"elevation"`
}

type Image struct {
	Url   string `json:"url"`
	Title string `json:"title"`
	Link  string `json:"link"`
}

type ObservationLocation struct {
	City      string `json:"city"`
	State     string `json:"state"`
	Country   string `json:"country"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Elevation string `json:"elevation"`
}
