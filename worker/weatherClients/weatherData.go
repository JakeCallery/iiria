package weatherClients

type WeatherData struct {
	Data struct {
		Timelines []struct {
			Timestep  string `json:"timestep"`
			StartTime string `json:"startTime"`
			EndTime   string `json:"endTime"`
			Intervals []struct {
				StartTime string `json:"startTime"`
				Values    struct {
					Temperature       float64 `json:"temperature"`
					PrecipitationType int     `json:"precipitationType"`
					WeatherCode       int     `json:"weatherCode"`
					UVIndex           int     `json:"uvIndex"`
					UVHealthConcern   int     `json:"uvHealthConcern"`
					Humidity          float64 `json:"humidity"`
				} `json:"values"`
			} `json:"intervals"`
		} `json:"timelines"`
	} `json:"data"`
}
