package models

import (
	"time"

	"github.com/google/uuid"
)

type File struct {
	ID          uint          `gorm:"primaryKey;autoIncrement" json:"id"`
	OwnerID     uuid.UUID     `gorm:"type:uuid;not null" json:"owner-id"`
	FileName    string        `gorm:"not null" json:"file-name"`
	FileType    string        `gorm:"not null" json:"file-type"`
	Processed   int           `gorm:"default:1" json:"processed"` // 1: Processing, 2: Processed, 3: Failed
	Affiliation []Affiliation `gorm:"foreignKey:FileID;references:ID;constraint:OnDelete:CASCADE" json:"affiliation"`
	CreatedAt   time.Time     `gorm:"autoCreateTime" json:"created-at"`
}
