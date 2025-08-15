package port

import (
	"github.com/EslamYasser-Dev/Go-AirScanner-Integration/core/entity"
	"github.com/google/uuid"
)

type OpreationRepository interface {
	NewOpreation(opreation *entity.Opreation) error
	GetOpreation(id uuid.UUID) (*entity.Opreation, error)
	GetAllOpreations() ([]*entity.Opreation, error)
	GetOpreationByCaseID(caseID uuid.UUID) ([]*entity.Opreation, error)
}
