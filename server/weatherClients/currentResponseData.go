package weatherClients

type CurrentResponseData struct {
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
				} `json:"values"`
			} `json:"intervals"`
		} `json:"timelines"`
	} `json:"data"`
}
