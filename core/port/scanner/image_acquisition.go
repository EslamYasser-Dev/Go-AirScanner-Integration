package scanner

import (
	"github.com/EslamYasser-Dev/Go-AirScanner-Integration/core/entity"
	"github.com/google/uuid"
)

type Scanner interface {
	FindAllDevices(deviceDispatcher *any) ([]*entity.Scanner, error)
	SelectScannerByID(id uuid.UUID) (*entity.Scanner, error)
	SelectScannerByDeviceID(deviceID string) (*entity.Scanner, error)
	ScanImages(deviceID, options *any) ([]*entity.DocumentData, error)
}
