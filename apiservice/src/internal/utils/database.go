package utils

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DbBaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey;column:id" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type DbDefaultSearchModel struct {
	Skip      int
	Limit     int
	OrderBy   string
	OrderType string
	Fields    []string
}
