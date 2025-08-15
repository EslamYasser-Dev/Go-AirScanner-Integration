package entity

import (
	"time"

	"github.com/google/uuid"
)

type Opreation struct {
	ID              uuid.UUID
	ScannerID       uuid.UUID
	UserID          uuid.UUID
	CaseID          uuid.UUID
	caseNo          string
	OperationResult string
	OperationError  string
	OperationStatus uint8
	OperationType   uint8
	TotalDocuments  int
	TimeStamp       time.Time
}

func (o *Opreation) GetCaseNo() string {
	return o.caseNo
}

func (o *Opreation) SetCaseNo(caseNo string) {
	//to do checks
	o.caseNo = caseNo
}
