package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/andreluzz/go-sql-builder/db"
	"github.com/cryo-management/api/config"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ServiceGroupTestSuite struct {
	suite.Suite
	InstanceID string
}

func (suite *ServiceGroupTestSuite) SetupTest() {
	pathSeparator := string(os.PathSeparator)
	file := fmt.Sprintf("..%sconfig.toml", pathSeparator)
	config := config.Config{}
	toml.DecodeFile(file, &config)
	db.Connect(config.Host, config.Port, config.User, config.Password, config.DBName, false)
}

func (suite *ServiceGroupTestSuite) Test00001CreateGroup() {
	data := map[string]interface{}{
		"name":        "Grupo Teste 01",
		"description": "Descrição do Grupo Teste 01",
		"code":        "grupoteste01",
		"active":      true,
	}
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "http://localhost:3333/api/v1/admin/groups", bytes.NewBuffer(jsonData))
	req.Header.Set("LanguageCode", "pt-br")

	response := CreateGroup(req)

	result := response.Data != nil && response.Code == 200
	groupValue := reflect.ValueOf(response.Data).Elem()
	suite.InstanceID = groupValue.FieldByName("ID").Interface().(string)

	assert.Equal(suite.T(), result, true)
}

func (suite *ServiceGroupTestSuite) Test00002LoadAllGroups() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/groups", nil)
	req.Header.Set("LanguageCode", "pt-br")

	response := LoadAllGroups(req)

	result := response.Data != nil && response.Code == 200

	assert.Equal(suite.T(), result, true)
}

func (suite *ServiceGroupTestSuite) Test00003LoadGroup() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/groups", nil)
	req.Header.Set("LanguageCode", "pt-br")

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("group_id", suite.InstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := LoadGroup(req)

	result := response.Data != nil && response.Code == 200

	assert.Equal(suite.T(), result, true)
}

func (suite *ServiceGroupTestSuite) Test00004UpdateGroup() {
	data := map[string]interface{}{
		"description": "Descrição do Grupo Teste 01 Updated",
		"active":      false,
	}
	jsonData, _ := json.Marshal(&data)

	req, _ := http.NewRequest("PATCH", "http://localhost:3333/api/v1/admin/groups", bytes.NewBuffer(jsonData))
	req.Header.Set("LanguageCode", "pt-br")

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("group_id", suite.InstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := UpdateGroup(req)

	result := response.Code == 200

	assert.Equal(suite.T(), result, true)
}

func (suite *ServiceGroupTestSuite) Test00005DeleteGroup() {
	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/groups", nil)
	req.Header.Set("LanguageCode", "pt-br")

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("group_id", suite.InstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := DeleteGroup(req)

	result := response.Code == 200

	assert.Equal(suite.T(), result, true)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestServiceGroupSuite(t *testing.T) {
	suite.Run(t, new(ServiceGroupTestSuite))
}
