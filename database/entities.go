package database

import (
	"time"
)

// Abstact entity, all other entities should be derived from it
type BaseEntity struct {
	ID uint64 `gorm:"primaryKey"`
}

type Migration struct {
	BaseEntity
	Version     string `gorm:"type:varchar(50);unique;not null"`
	Description string `gorm:"type:varchar(256)"`
	ExecutedAt  time.Time
	Duration    int
	Status      MigrationStatus `gorm:"type:varchar(20)"`
}
