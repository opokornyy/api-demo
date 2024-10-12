package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"

	"api-demo/internal/controller/schema"
	"api-demo/internal/model"
)

type UserRepository interface {
	CreateUser(*model.User) error
	GetUser(uuid.UUID) (*model.User, error)
}

type UserController struct {
	repo     UserRepository
	validate *validator.Validate
}

func NewUserController(repo UserRepository, validate *validator.Validate) *UserController {
	return &UserController{repo, validate}
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	req, err := parseCreateUserBody(r, uc.validate)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Convert request to user model
	user, err := convertRequestToUser(req)
	if err != nil {
		log.Error().Err(err).Msgf("failed to parse request: %v", err)
		respondWithError(w, http.StatusInternalServerError, "failed to parse request to user")
		return
	}

	// Create user
	if err := uc.repo.CreateUser(user); err != nil {
		log.Error().Err(err).Msgf("failed to create user: %v", err)
		respondWithError(w, http.StatusInternalServerError, "failed to create user")
		return
	}

	log.Info().Msgf("User created: %+v", req)
	respondWithJSON(w, http.StatusCreated, user)
}

func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	// Parse id to uuid
	uuid, err := uuid.Parse(id)
	if err != nil {
		log.Error().Err(err).Msgf("failed to parse id: %v", err)
		respondWithError(w, http.StatusBadRequest, "invalid id")
		return
	}

	// Get user
	user, err := uc.repo.GetUser(uuid)
	if err != nil {
		log.Error().Err(err).Msgf("failed to get user: %v", err)
		respondWithError(w, http.StatusNotFound, "user not found")
		return
	}

	log.Info().Msgf("User retrieved: %+v", user)
	respondWithJSON(w, http.StatusOK, user)
}

func parseCreateUserBody(r *http.Request, validate *validator.Validate) (*schema.CreateUserRequest, error) {
	var req schema.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Error().Err(err).Msgf("invalid json body: %v", err)
		return nil, fmt.Errorf("error decoding request: %v", err)
	}

	// Validate request
	if err := validate.Struct(req); err != nil {
		log.Error().Err(err).Msgf("validation error: %v", err)
		return nil, fmt.Errorf("error validating request: %v", err)
	}

	return &req, nil
}

func convertRequestToUser(req *schema.CreateUserRequest) (*model.User, error) {
	// Parse DateOfBirth from string to time.Time
	parsedTime, err := time.Parse(time.RFC3339, req.DateOfBirth)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:          req.ID,
		Name:        req.Name,
		Email:       req.Email,
		DateOfBirth: parsedTime,
	}, nil
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
