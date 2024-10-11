package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name        string    `json:"name"`
	Email       string    `gorm:"unique;not null" json:"email"`
	DateOfBirth time.Time `json:"dateOfBirth"`
}
