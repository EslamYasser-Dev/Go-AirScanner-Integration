package scanner

import (
	"context"

	"github.com/EslamYasser-Dev/Go-AirScanner-Integration/core/entity"
	"github.com/go-ole/go-ole"
)

type Do interface {
	ListScanners(ctx context.Context, devManager *ole.IUnknown) ([]*entity.Scanner, error)
	Scan(ctx context.Context, devManager *ole.IUnknown) ([]*entity.DocumentData, error)
}
