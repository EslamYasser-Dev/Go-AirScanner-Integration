package entity

import (
	"time"

	"github.com/google/uuid"
)

// DocumentMetaData holds metadata about a scanned document.
// GORM tags have been added for database mapping.
// ID == file name without extension
type DocumentMetaData struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	CaseNumber   string    `gorm:"type:uuid" json:"case_number"`
	UserID       uuid.UUID `gorm:"type:uuid" json:"user_id"`
	ScannerID    uuid.UUID `gorm:"type:uuid" json:"scanner_id"`
	DocumentType uint8     `gorm:"type:tinyint" json:"document_type"` //location on disk, cloud, s3 etc.
	FileName     string    `gorm:"type:varchar(255);not null" json:"document_name"`
	Url          string    `gorm:"type:varchar(255)" json:"url"`
	Timestamp    time.Time `json:"timestamp"`
}

// DocumentData holds the raw data of the document.
type DocumentData []byte
