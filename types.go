package main

import (
	"time"
)

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
