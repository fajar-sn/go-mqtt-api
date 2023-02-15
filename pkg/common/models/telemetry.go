package models

import "time"

type Telemetry struct {
	ID			uint		`gorm:"primaryKey" json:"id"`
	Timestamp	time.Time	`json:"timestamp"`
	Data		int			`json:"data"`
	DeviceID	uint		`json:"device_id"`
	Device		Device		`json:"device"`
}

type TelemetryRequestBody struct {
	Data		int		`json:"data"`
	Token		string 	`json:"token"`
}