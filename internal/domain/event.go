package domain

import "time"

type Event struct {
	Timestamp          time.Time
	SensorSerialNumber string
	SensorID           int64
	Payload            int64
}
