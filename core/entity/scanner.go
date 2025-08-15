package entity

import "github.com/google/uuid"

type ScannerStatus string

const (
	ScannerStatusAvailable ScannerStatus = "Available"
	ScannerStatusBusy      ScannerStatus = "Busy"
	ScannerStatusOffline   ScannerStatus = "Offline"
)

type Scanner struct {
	ID       uuid.UUID     `json:"id"`
	Name     string        `json:"name"`
	DeviceID string        `json:"device_id"`
	Status   ScannerStatus `json:"status"`
}

func NewScanner(name, deviceID string) *Scanner {
	return &Scanner{
		ID:       uuid.New(),
		Name:     name,
		DeviceID: deviceID,
		Status:   ScannerStatusOffline,
	}
}
