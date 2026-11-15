package models

import "gorm.io/gorm"

type IMG struct {
	gorm.Model
	Name string `gorm:"not null"`
	Type string `gorm:"not null"`
	Data []byte `gorm:"type:bytea;not null"`
}
