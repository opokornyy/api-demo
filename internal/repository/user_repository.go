package repository

// TODO database interface

type UserRepository struct {
	// TODO add database connection
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) CreateUser() {
	// TODO implement
}

func (r *UserRepository) GetUser() {
	// TODO implement
}
