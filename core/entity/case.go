package entity

import (
	"time"

	"github.com/google/uuid"
)

type CaseStatus string

const (
	CaseStatusDone     CaseStatus = "done"
	CaseStatusRejected CaseStatus = "rejected"
	CaseStatusPending  CaseStatus = "pending"
)

type Case struct {
	ID         uuid.UUID  `gorm:"type:uuid;primary_key;" json:"id"`
	CaseNumber string     `gorm:"type:varchar(100);unique_index" json:"case_number"`
	UserID     uuid.UUID  `gorm:"type:uuid" json:"user_id"`
	Timestamp  time.Time  `json:"timestamp"`
	Status     CaseStatus `gorm:"type:varchar(20)" json:"status"`
}

func NewCase(caseNumber string, userID uuid.UUID) *Case {
	return &Case{
		ID:         uuid.New(),
		CaseNumber: caseNumber,
		UserID:     userID,
		Timestamp:  time.Now(),
		Status:     CaseStatusPending,
	}
}
