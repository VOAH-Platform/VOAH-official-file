package models

import (
	"github.com/google/uuid"
	"implude.kr/VOAH-Official-File/configs"
)

type Affiliation struct {
	ID                uuid.UUID          `gorm:"type:uuid;primaryKey" json:"id"`
	AffiliationType   configs.ObjectType `gorm:"not null" json:"affiliation-type"`
	AffiliationTarget uuid.UUID          `gorm:"type:uuid;not null" json:"affiliation-target"`
	FileID            uint               `gorm:"not null" json:"file-id"`
}
