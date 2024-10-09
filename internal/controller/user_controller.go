package controller

type UserRepository interface {
	CreateUser()
	GetUser()
}

type UserController struct {
	repo UserRepository
}

func NewUserController(repo UserRepository) *UserController {
	return &UserController{repo: repo}
}

func (uc *UserController) CreateUser() {
	// TODO: implement
}

func (uc *UserController) GetUser() {
	// TODO: implement
}
