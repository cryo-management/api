package services

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	"github.com/andreluzz/go-sql-builder/db"
	"github.com/cryo-management/api/config"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ServiceUserTestSuite struct {
	suite.Suite
	InstanceID string
	UserID     string
}

func (suite *ServiceUserTestSuite) SetupTest() {
	config, _ := config.NewConfig("..\\config.toml")
	db.Connect(config.Host, config.Port, config.User, config.Password, config.DBName, false)
	suite.UserID = "57a97aaf-16da-44ef-a8be-b1caf52becd6"
}

func (suite *ServiceUserTestSuite) Test00001CreateUser() {
	data := map[string]interface{}{
		"username":      "usuarioteste01",
		"first_name":    "Usuário",
		"last_name":     "Teste 01",
		"email":         "usuarioteste01@domain.com",
		"password":      "123456",
		"language_code": "pt-br",
		"active":        true,
	}
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "http://localhost:3333/api/v1/admin/users", bytes.NewBuffer(jsonData))
	req.Header.Set("languageCode", "pt-br")
	req.Header.Set("userID", suite.UserID)

	response := CreateUser(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)

	userValue := reflect.ValueOf(response.Data).Elem()
	suite.InstanceID = userValue.FieldByName("ID").Interface().(string)
}

func (suite *ServiceUserTestSuite) Test00002LoadAllUsers() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/users", nil)
	req.Header.Set("languageCode", "pt-br")
	req.Header.Set("userID", suite.UserID)

	response := LoadAllUsers(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceUserTestSuite) Test00003LoadUser() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/users", nil)
	req.Header.Set("languageCode", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("user_id", suite.InstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := LoadUser(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceUserTestSuite) Test00004UpdateUser() {
	data := map[string]interface{}{
		"last_name": "User 01 Updated",
		"active":    false,
	}
	jsonData, _ := json.Marshal(&data)

	req, _ := http.NewRequest("PATCH", "http://localhost:3333/api/v1/admin/users", bytes.NewBuffer(jsonData))
	req.Header.Set("languageCode", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("user_id", suite.InstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := UpdateUser(req)

	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceUserTestSuite) Test00005DeleteUser() {
	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/users", nil)
	req.Header.Set("languageCode", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("user_id", suite.InstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := DeleteUser(req)

	assert.Equal(suite.T(), 200, response.Code)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestServiceUserSuite(t *testing.T) {
	suite.Run(t, new(ServiceUserTestSuite))
}
