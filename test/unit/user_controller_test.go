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

// MockUserRepository
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

const expectedUser = `{"id":"93789a5f-5ff1-49a9-8e5b-601879bbe522","name":"John Doe","email":"john@doe.com","dateOfBirth":"2000-01-01T00:00:00Z"}`

func TestCreateUser(t *testing.T) {
	reqBody := schema.CreateUserRequest{
		ID:          uuid.MustParse("93789a5f-5ff1-49a9-8e5b-601879bbe522"),
		Name:        "John Doe",
		Email:       "john@doe.com",
		DateOfBirth: "2000-01-01T00:00:00Z",
	}

	reqBodyJSON, err := json.Marshal(reqBody)
	assert.NoError(t, err)

	// Create test request
	req, err := http.NewRequest("POST", "/save", bytes.NewBuffer(reqBodyJSON))
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

	// Assert the status code and response body
	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Equal(t, expectedUser, rr.Body.String())
}
