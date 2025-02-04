package integration_test

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"api-demo/internal/app"

	"github.com/stretchr/testify/assert"
)

const (
	serverURL = "http://localhost:8080/"

	testUUID        = "86f8cebb-7cfb-41b7-9c2c-84cfc3a3dda2"
	notExistingUUID = "86f8cebb-7cfb-41b7-9c2c-84cfc3a3dda3"

	// Missing email field
	invalidUserData = `{
		"id": "86f8cebb-7cfb-41b7-9c2c-84cfc3a3dda2",
		"name": "Ho Lee Fock",
		"date_of_birth": "2020-01-01T12:12:34+00:00"
	}`

	userData = `{
		"id": "86f8cebb-7cfb-41b7-9c2c-84cfc3a3dda2",
		"name": "Ho Lee Fock",
		"email": "holee@fock.com",
		"date_of_birth": "2020-01-01T12:12:34+00:00"
	}`
	expectedUserData = `{"id":"86f8cebb-7cfb-41b7-9c2c-84cfc3a3dda2","name":"Ho Lee Fock","email":"holee@fock.com","dateOfBirth":"2020-01-01T12:12:34Z"}`
)

func TestMain(m *testing.M) {
	// Start the server
	go func() {
		app.Run()
	}()

	// Give the server some time to start up
	time.Sleep(2 * time.Second)

	// Run tests
	exitVal := m.Run()

	// Exit with the exit value from the tests
	os.Exit(exitVal)
}

func Test_CreateAndRetrieveUser_Succeeds(t *testing.T) {
	// Create a new user
	resp, err := http.Post(serverURL+"save", "application/json", strings.NewReader(userData))
	assert.NoError(t, err)
	defer resp.Body.Close()

	// Assert that the user was created
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	// Retrieve the user you just created
	resp, err = http.Get(fmt.Sprintf(serverURL+"%s", testUUID))
	assert.NoError(t, err)
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	// Assert that the correct user was retrieved
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, expectedUserData, string(body))
}

func Test_CreateUser_InvalidRequest_Returns_BadRequest(t *testing.T) {
	// Create a new user with invalid data
	resp, err := http.Post(serverURL+"save", "application/json", strings.NewReader(invalidUserData))
	assert.NoError(t, err)
	defer resp.Body.Close()

	// Assert that the user was not created
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func Test_GetUser_NonExisting_User_Returns_NotFound(t *testing.T) {
	// Retrieve non-existing user
	resp, err := http.Get(fmt.Sprintf(serverURL+"%s", notExistingUUID))
	assert.NoError(t, err)
	defer resp.Body.Close()

	// Assert that the user was not found
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func Test_GetUser_InvalidUUID_Returns_BadRequest(t *testing.T) {
	// Retrieve non-existing user
	resp, err := http.Get(serverURL + "invalid-uuid")
	assert.NoError(t, err)
	defer resp.Body.Close()

	// Assert that bad request was returned
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}
