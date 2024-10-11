package schema

import "github.com/google/uuid"

type CreateUserRequest struct {
	ID          uuid.UUID `validate:"required,uuid4" json:"id"`
	Name        string    `validate:"required,min=1,max=200" json:"name"`
	Email       string    `validate:"required,email" json:"email"`
	DateOfBirth string    `validate:"required,datetime=2006-01-02T15:04:05Z07:00" json:"date_of_birth"`
}
