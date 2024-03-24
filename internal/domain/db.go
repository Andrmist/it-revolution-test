package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Link struct {
	ID           string `gorm:"primaryKey" json:"-"`
	OriginalLink string `json:"original_link"`
	Count        int    `gorm:"default:0" json:"count"`
}

func (l *Link) BeforeCreate(db *gorm.DB) (err error) {
	l.ID = uuid.NewString()[:4]
	var exists Link
	db.Where("id = ?", l.ID).First(&exists)
	if exists.ID != "" {
		return l.BeforeCreate(db)
	}
	return
}
