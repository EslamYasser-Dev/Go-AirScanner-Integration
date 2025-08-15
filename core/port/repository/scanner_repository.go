package port

import (
	"github.com/EslamYasser-Dev/Go-AirScanner-Integration/core/entity"
	"github.com/google/uuid"
)

type ScannerRepository interface {
	FindAll() ([]*entity.Scanner, error)
	SelectScannerByID(id uuid.UUID) (*entity.Scanner, error)
	SelectScannerByDeviceID(deviceID string) (*entity.Scanner, error)
	ScanImages(deviceID string, options any) ([]*entity.DocumentData, error)
}
