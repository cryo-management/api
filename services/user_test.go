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
}

func (suite *ServiceUserTestSuite) SetupTest() {
	config, _ := config.NewConfig("..\\config.toml")
	db.Connect(config.Host, config.Port, config.User, config.Password, config.DBName, false)
}

func (suite *ServiceUserTestSuite) Test00001CreateUser() {
	data := map[string]interface{}{
		"first_name": "Usuário",
		"last_name":  "Teste 01",
		"email":      "usuarioteste01@domain.com",
		"password":   "123456",
		"language":   "pt-br",
		"active":     true,
	}
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "http://localhost:3333/api/v1/admin/users", bytes.NewBuffer(jsonData))
	req.Header.Set("LanguageCode", "pt-br")

	response := CreateUser(req)

	result := response.Data != nil && response.Code == 200
	userValue := reflect.ValueOf(response.Data).Elem()
	suite.InstanceID = userValue.FieldByName("ID").Interface().(string)

	assert.Equal(suite.T(), result, true)
}

func (suite *ServiceUserTestSuite) Test00002LoadAllUsers() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/users", nil)
	req.Header.Set("LanguageCode", "pt-br")

	response := LoadAllUsers(req)

	result := response.Data != nil && response.Code == 200

	assert.Equal(suite.T(), result, true)
}

func (suite *ServiceUserTestSuite) Test00003LoadUser() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/users", nil)
	req.Header.Set("LanguageCode", "pt-br")

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("user_id", suite.InstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := LoadUser(req)

	result := response.Data != nil && response.Code == 200

	assert.Equal(suite.T(), result, true)
}

func (suite *ServiceUserTestSuite) Test00004UpdateUser() {
	data := map[string]interface{}{
		"last_name": "User 01 Updated",
		"active":    false,
	}
	jsonData, _ := json.Marshal(&data)

	req, _ := http.NewRequest("PATCH", "http://localhost:3333/api/v1/admin/users", bytes.NewBuffer(jsonData))
	req.Header.Set("LanguageCode", "pt-br")

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("user_id", suite.InstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := UpdateUser(req)

	result := response.Code == 200

	assert.Equal(suite.T(), result, true)
}

func (suite *ServiceUserTestSuite) Test00005DeleteUser() {
	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/users", nil)
	req.Header.Set("LanguageCode", "pt-br")

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("user_id", suite.InstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := DeleteUser(req)

	result := response.Code == 200

	assert.Equal(suite.T(), result, true)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestServiceUserSuite(t *testing.T) {
	suite.Run(t, new(ServiceUserTestSuite))
}
