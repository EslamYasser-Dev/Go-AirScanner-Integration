package repository

import (
	"context"

	"github.com/EslamYasser-Dev/Go-AirScanner-Integration/core/entity"
	"github.com/google/uuid"
)

type Case interface {
	OpenNewCase(ctx context.Context, caseNumber string, userID uuid.UUID) (entity.Case, error)
	SyncPendingCases(ctx context.Context) error
	GetCasesBasedOnStatus(ctx context.Context, status entity.CaseStatus) ([]*entity.Case, error)
	
}
