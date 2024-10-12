package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"api-demo/internal/controller"
	"api-demo/internal/controller/schema"
	"api-demo/internal/model"
)

// MockUserRepository is a mock implementation of UserRepository for testing
type MockUserRepository struct {
	user model.User
	err  error
}

func NewMockUserRepository(user model.User, err error) *MockUserRepository {
	return &MockUserRepository{user, err}
}

// CreateUser simulates creating a user
func (m *MockUserRepository) CreateUser(user *model.User) error {
	return m.err
}

// GetUser simulates retrieving a user by ID
func (m *MockUserRepository) GetUser(id uuid.UUID) (*model.User, error) {
	return &m.user, m.err
}

func TestCreateUser(t *testing.T) {
	// TODO: maybe move this to a helper function
	// or even a json file
	reqBody := schema.CreateUserRequest{
		ID:          uuid.New(),
		Name:        "John Doe",
		Email:       "john@doe.com",
		DateOfBirth: "2000-01-01T00:00:00Z", // Use ISO 8601 format for date
	}

	reqBodyJSON, err := json.Marshal(reqBody)
	assert.NoError(t, err)

	// Create test request
	req, err := http.NewRequest("POST", "/user", bytes.NewBuffer(reqBodyJSON))
	if err != nil {
		t.Fatal(err)
	}

	mockUserRepository := NewMockUserRepository(model.User{}, nil)
	validator := validator.New(validator.WithRequiredStructEnabled())

	// Create a UserController with a MockUserRepository and a validator
	uc := controller.NewUserController(mockUserRepository, validator)

	// Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(uc.CreateUser)

	// Call the handler with the request and response recorder.
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// Check the response body is what we expect.
	// TODO: rename the reqBodyJSON to expected
	// expected := `{"id":"26908e04-868c-4d8e-85f8-6b1284dcf750","name":"Mike Oxlong","email":"mike@oxlong.cz","date_of_birth":"2020-01-01T12:12:34+00:00"}`
	if rr.Body.String() != string(reqBodyJSON) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), reqBodyJSON)
	}
}
