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

type ServiceSchemaTestSuite struct {
	suite.Suite
	InstanceID string
}

func (suite *ServiceSchemaTestSuite) SetupTest() {
	pathSeparator := string(os.PathSeparator)
	file := fmt.Sprintf("..%sconfig.toml", pathSeparator)
	config := config.Config{}
	toml.DecodeFile(file, &config)
	db.Connect(config.Host, config.Port, config.User, config.Password, config.DBName, false)
}

func (suite *ServiceSchemaTestSuite) Test00001CreateSchema() {
	data := map[string]interface{}{
		"name":        "Schema Teste 01",
		"description": "Descrição do Schema Teste 01",
		"code":        "schemateste01",
		"module":      false,
		"active":      true,
	}
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "http://localhost:3333/api/v1/admin/schemas", bytes.NewBuffer(jsonData))
	req.Header.Set("LanguageCode", "pt-br")

	response := CreateSchema(req)

	result := response.Data != nil && response.Code == 200
	schemaValue := reflect.ValueOf(response.Data).Elem()
	suite.InstanceID = schemaValue.FieldByName("ID").Interface().(string)

	assert.Equal(suite.T(), result, true)
}

func (suite *ServiceSchemaTestSuite) Test00002LoadAllSchemas() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/schemas", nil)
	req.Header.Set("LanguageCode", "pt-br")

	response := LoadAllSchemas(req)

	result := response.Data != nil && response.Code == 200

	assert.Equal(suite.T(), result, true)
}

func (suite *ServiceSchemaTestSuite) Test00003LoadSchema() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/schemas", nil)
	req.Header.Set("LanguageCode", "pt-br")

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("schema_id", suite.InstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := LoadSchema(req)

	result := response.Data != nil && response.Code == 200

	assert.Equal(suite.T(), result, true)
}

func (suite *ServiceSchemaTestSuite) Test00004UpdateSchema() {
	data := map[string]interface{}{
		"description": "Descrição do Schema Teste 01 Updated",
		"active":      false,
	}
	jsonData, _ := json.Marshal(&data)

	req, _ := http.NewRequest("PATCH", "http://localhost:3333/api/v1/admin/schemas", bytes.NewBuffer(jsonData))
	req.Header.Set("LanguageCode", "pt-br")

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("schema_id", suite.InstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := UpdateSchema(req)

	result := response.Code == 200

	assert.Equal(suite.T(), result, true)
}

func (suite *ServiceSchemaTestSuite) Test00005DeleteSchema() {
	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/schemas", nil)
	req.Header.Set("LanguageCode", "pt-br")

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("schema_id", suite.InstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := DeleteSchema(req)

	result := response.Code == 200

	assert.Equal(suite.T(), result, true)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestServiceSchemaSuite(t *testing.T) {
	suite.Run(t, new(ServiceSchemaTestSuite))
}
