package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	DateOfBirth time.Time `json:"dateOfBirth"` // Use time.Time for date fields?
}
