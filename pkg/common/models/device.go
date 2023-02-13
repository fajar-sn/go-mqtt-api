package models

import (
	"time"
)

type Device struct {
	ID			uint		`gorm:"primaryKey" json:"id"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
	Name		string		`json:"name"`
	Token 		string 		`json:"token"`
}

type DeviceRequestBody struct {
	Name	string `json:"name"`
	Token	string `json:"token"`
}