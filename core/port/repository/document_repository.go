package repository

import (
	"context"

	"github.com/EslamYasser-Dev/Go-AirScanner-Integration/core/entity"
	"github.com/google/uuid"
)

type Document interface {
	Save(ctx context.Context, document *entity.DocumentMetaData, data []byte) error
	Find(ctx context.Context, id uuid.UUID) (*entity.DocumentMetaData, error)
	FindByCaseNumber(ctx context.Context, caseNumber string) ([]*entity.DocumentMetaData, error)
	PrepareForSync(ctx context.Context) error
}
