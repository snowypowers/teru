package store

import (
	"time"
	"encoding/json"
	"log"
	"strings"
)

var areaNames = []string{
"Ang Mo Kio",
"Bedok",
"Bishan",
"Boon Lay",
"Bukit Batok",
"Bukit Merah",
"Bukit Panjang",
"Bukit Timah",
"Central Water Catchment",
"Changi",
"Choa Chu Kang",
"Clementi",
"City",
"Geylang",
"Hougang",
"Jalan Bahar",
"Jurong East",
"Jurong Island",
"Jurong West",
"Kallang",
"Lim Chu Kang",
"Mandai",
"Marine Parade",
"Novena",
"Pasir Ris",
"Paya Lebar",
"Pioneer",
"Pulau Tekong",
"Pulau Ubin",
"Punggol",
"Queenstown",
"Seletar",
"Sembawang",
"Sengkang",
"Sentosa",
"Serangoon",
"Southern Islands",
"Sungei Kadut",
"Tampines",
"Tanglin",
"Tengah",
"Toa Payoh",
"Tuas",
"Western Islands",
"Western Water Catchment",
"Woodlands",
"Yishun"}

//WF2 Raw data struct
type WF2 struct {
	AreaMetadata []struct {
		Name string `json:"name"`
		LabelLocation struct {
			Latitude float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"label_location"`
	} `json:"area_metadata"`
	Items []struct {
		UpdateTimestamp time.Time `json:"update_timestamp"`
		Timestamp time.Time `json:"timestamp"`
		ValidPeriod struct {
			Start time.Time `json:"start"`
			End time.Time `json:"end"`
		} `json:"valid_period"`
		Forecasts []struct {
			Area string `json:"area"`
			Forecast string `json:"forecast"`
		} `json:"forecasts"`
	} `json:"items"`
	APIInfo struct {
		Status string `json:"status"`
	} `json:"api_info"`
}

type WF2Update struct {
	Timestamp time.Time
	Forecasts map[string]string
}

//ParseWf2 into a WF2Update
func ParseWf2(data []byte) WF2Update {
	var record WF2
	if err := json.Unmarshal(data, &record); err != nil {
		log.Println(err)
	}
	ts := record.Items[0].Timestamp
	forecasts := make(map[string]string)
	update := WF2Update{ts, forecasts}
	for _,i := range record.Items[0].Forecasts {
		forecasts[i.Area] = i.Forecast
	}
	return update
}

//ParseArea parses string into a possible Area.
//Returns empty string if no matches
func ParseArea(args string) string {
	sanitised := strings.Trim(args, " .")
	sanitised = strings.Title(sanitised)
	for _,i:= range areaNames {
		if sanitised == i {
			return i
		}
	}
	return ""
}
