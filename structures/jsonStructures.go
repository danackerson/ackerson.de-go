package structures

var TestGeoLocationPost string = `{"jsonrpc":"2.0","method":"post","params":{"lat":48.3003496,"lng":11.356716099999971},"id":14}`

type JsonGeoLocationRequest struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Lat float32 `json:"lat"`
		Lng float32 `json:"lng"`
	} `json:"params"`
	Id int `json:"id"`
}

type CurrentWeatherForecast struct {
	Response struct {
		Version        string `json:"version"`
		TermsofService string `json:"termsofService"`
		Features       struct {
			Forecast int `json:"forecast"`
		} `json:"features"`
	} `json:"response"`
	Forecast struct {
		TxtForecast struct {
			Date        string `json:"date"`
			Forecastday []struct {
				Period        int    `json:"period"`
				Icon          string `json:"icon"`
				IconUrl       string `json:"icon_url"`
				Title         string `json:"title"`
				Fcttext       string `json:"fcttext"`
				FcttextMetric string `json:"fcttext_metric"`
				Pop           string `json:"pop"`
			} `json:"forecastday"`
		} `json:"txt_forecast"`
		Simpleforecast struct {
			Forecastday []struct {
				Date struct {
					Epoch          string `json:"epoch"`
					Pretty         string `json:"pretty"`
					Day            int    `json:"day"`
					Month          int    `json:"month"`
					Year           int    `json:"year"`
					Yday           int    `json:"yday"`
					Hour           int    `json:"hour"`
					Min            string `json:"min"`
					Sec            int    `json:"sec"`
					Isdst          string `json:"isdst"`
					Monthname      string `json:"monthname"`
					MonthnameShort string `json:"monthname_short"`
					WeekdayShort   string `json:"weekday_short"`
					Weekday        string `json:"weekday"`
					Ampm           string `json:"ampm"`
					TzShort        string `json:"tz_short"`
					TzLong         string `json:"tz_long"`
				} `json:"date"`
				Period int `json:"period"`
				High   struct {
					Fahrenheit string `json:"fahrenheit"`
					Celsius    string `json:"celsius"`
				} `json:"high"`
				Low struct {
					Fahrenheit string `json:"fahrenheit"`
					Celsius    string `json:"celsius"`
				} `json:"low"`
				Conditions string `json:"conditions"`
				Icon       string `json:"icon"`
				IconUrl    string `json:"icon_url"`
				Skyicon    string `json:"skyicon"`
				Pop        int    `json:"pop"`
				QpfAllday  struct {
					In int `json:"in"`
					Mm int `json:"mm"`
				} `json:"qpf_allday"`
				QpfDay struct {
					In interface{} `json:"in"`
					Mm interface{} `json:"mm"`
				} `json:"qpf_day"`
				QpfNight struct {
					In int `json:"in"`
					Mm int `json:"mm"`
				} `json:"qpf_night"`
				SnowAllday struct {
					In int `json:"in"`
					Cm int `json:"cm"`
				} `json:"snow_allday"`
				SnowDay struct {
					In interface{} `json:"in"`
					Cm interface{} `json:"cm"`
				} `json:"snow_day"`
				SnowNight struct {
					In int `json:"in"`
					Cm int `json:"cm"`
				} `json:"snow_night"`
				Maxwind struct {
					Mph     int    `json:"mph"`
					Kph     int    `json:"kph"`
					Dir     string `json:"dir"`
					Degrees int    `json:"degrees"`
				} `json:"maxwind"`
				Avewind struct {
					Mph     int    `json:"mph"`
					Kph     int    `json:"kph"`
					Dir     string `json:"dir"`
					Degrees int    `json:"degrees"`
				} `json:"avewind"`
				Avehumidity int `json:"avehumidity"`
				Maxhumidity int `json:"maxhumidity"`
				Minhumidity int `json:"minhumidity"`
			} `json:"forecastday"`
		} `json:"simpleforecast"`
	} `json:"forecast"`
}

type CurrentWeatherConditions struct {
	Response struct {
		Version        string `json:"version"`
		TermsofService string `json:"termsofService"`
		Features       struct {
			Conditions int `json:"conditions"`
		} `json:"features"`
	} `json:"response"`
	CurrentObservation struct {
		Image struct {
			Url   string `json:"url"`
			Title string `json:"title"`
			Link  string `json:"link"`
		} `json:"image"`
		DisplayLocation struct {
			Full           string `json:"full"`
			City           string `json:"city"`
			State          string `json:"state"`
			StateName      string `json:"state_name"`
			Country        string `json:"country"`
			CountryIso3166 string `json:"country_iso3166"`
			Zip            string `json:"zip"`
			Magic          string `json:"magic"`
			Wmo            string `json:"wmo"`
			Latitude       string `json:"latitude"`
			Longitude      string `json:"longitude"`
			Elevation      string `json:"elevation"`
		} `json:"display_location"`
		ObservationLocation struct {
			Full           string `json:"full"`
			City           string `json:"city"`
			State          string `json:"state"`
			Country        string `json:"country"`
			CountryIso3166 string `json:"country_iso3166"`
			Latitude       string `json:"latitude"`
			Longitude      string `json:"longitude"`
			Elevation      string `json:"elevation"`
		} `json:"observation_location"`
		Estimated struct {
		} `json:"estimated"`
		StationId             string  `json:"station_id"`
		ObservationTime       string  `json:"observation_time"`
		ObservationTimeRfc822 string  `json:"observation_time_rfc822"`
		ObservationEpoch      string  `json:"observation_epoch"`
		LocalTimeRfc822       string  `json:"local_time_rfc822"`
		LocalEpoch            string  `json:"local_epoch"`
		LocalTzShort          string  `json:"local_tz_short"`
		LocalTzLong           string  `json:"local_tz_long"`
		LocalTzOffset         string  `json:"local_tz_offset"`
		Weather               string  `json:"weather"`
		TemperatureString     string  `json:"temperature_string"`
		TempF                 int     `json:"temp_f"`
		TempC                 float32 `json:"temp_c"`
		RelativeHumidity      string  `json:"relative_humidity"`
		WindString            string  `json:"wind_string"`
		WindDir               string  `json:"wind_dir"`
		WindDegrees           int     `json:"wind_degrees"`
		WindMph               int     `json:"wind_mph"`
		WindGustMph           string  `json:"wind_gust_mph"`
		WindKph               float32 `json:"wind_kph"`
		WindGustKph           string  `json:"wind_gust_kph"`
		PressureMb            string  `json:"pressure_mb"`
		PressureIn            string  `json:"pressure_in"`
		PressureTrend         string  `json:"pressure_trend"`
		DewpointString        string  `json:"dewpoint_string"`
		DewpointF             int     `json:"dewpoint_f"`
		DewpointC             int     `json:"dewpoint_c"`
		HeatIndexString       string  `json:"heat_index_string"`
		HeatIndexF            string  `json:"heat_index_f"`
		HeatIndexC            string  `json:"heat_index_c"`
		WindchillString       string  `json:"windchill_string"`
		WindchillF            string  `json:"windchill_f"`
		WindchillC            string  `json:"windchill_c"`
		FeelslikeString       string  `json:"feelslike_string"`
		FeelslikeF            string  `json:"feelslike_f"`
		FeelslikeC            string  `json:"feelslike_c"`
		VisibilityMi          string  `json:"visibility_mi"`
		VisibilityKm          string  `json:"visibility_km"`
		Solarradiation        string  `json:"solarradiation"`
		UV                    string  `json:"UV"`
		Precip1hrString       string  `json:"precip_1hr_string"`
		Precip1hrIn           string  `json:"precip_1hr_in"`
		Precip1hrMetric       string  `json:"precip_1hr_metric"`
		PrecipTodayString     string  `json:"precip_today_string"`
		PrecipTodayIn         string  `json:"precip_today_in"`
		PrecipTodayMetric     string  `json:"precip_today_metric"`
		Icon                  string  `json:"icon"`
		IconUrl               string  `json:"icon_url"`
		ForecastUrl           string  `json:"forecast_url"`
		HistoryUrl            string  `json:"history_url"`
		ObUrl                 string  `json:"ob_url"`
		Nowcast               string  `json:"nowcast"`
	} `json:"current_observation"`
}