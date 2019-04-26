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

type ServiceSchemaTestSuite struct {
	suite.Suite
	SchemaInstanceID string
	ModuleInstanceID string
	UserID           string
}

func (suite *ServiceSchemaTestSuite) SetupTest() {
	config, _ := config.NewConfig("..\\config.toml")
	db.Connect(config.Host, config.Port, config.User, config.Password, config.DBName, false)
	suite.UserID = "57a97aaf-16da-44ef-a8be-b1caf52becd6"
}

func (suite *ServiceSchemaTestSuite) Test00001CreateSchema() {
	createModuleToSchema(suite)

	data := map[string]interface{}{
		"name":        "Schema Teste 01",
		"description": "Descrição do Schema Teste 01",
		"code":        "schemateste01",
		"module":      false,
		"active":      true,
	}
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "http://localhost:3333/api/v1/admin/schemas", bytes.NewBuffer(jsonData))
	req.Header.Set("languageCode", "pt-br")
	req.Header.Set("userID", suite.UserID)

	response := CreateSchema(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)

	schemaValue := reflect.ValueOf(response.Data).Elem()
	suite.SchemaInstanceID = schemaValue.FieldByName("ID").Interface().(string)
}

func (suite *ServiceSchemaTestSuite) Test00002LoadAllSchemas() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/schemas", nil)
	req.Header.Set("languageCode", "pt-br")
	req.Header.Set("userID", suite.UserID)

	response := LoadAllSchemas(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceSchemaTestSuite) Test00003LoadSchema() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/schemas", nil)
	req.Header.Set("languageCode", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("schema_id", suite.SchemaInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := LoadSchema(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceSchemaTestSuite) Test00004UpdateSchema() {
	data := map[string]interface{}{
		"description": "Descrição do Schema Teste 01 Updated",
		"active":      false,
	}
	jsonData, _ := json.Marshal(&data)

	req, _ := http.NewRequest("PATCH", "http://localhost:3333/api/v1/admin/schemas", bytes.NewBuffer(jsonData))
	req.Header.Set("languageCode", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("schema_id", suite.SchemaInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := UpdateSchema(req)

	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceSchemaTestSuite) Test00005InsertModuleInSchema() {
	req, _ := http.NewRequest("POST", "http://localhost:3333/api/v1/admin/schemas", nil)
	req.Header.Set("languageCode", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("schema_id", suite.SchemaInstanceID)
	rctx.URLParams.Add("module_id", suite.ModuleInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := InsertModuleInSchema(req)

	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceSchemaTestSuite) Test00006LoadAllModulesBySchema() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/schemas", nil)
	req.Header.Set("languageCode", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("schema_id", suite.SchemaInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := LoadAllModulesBySchema(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceSchemaTestSuite) Test00007RemoveModuleFromSchema() {
	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/schemas", nil)
	req.Header.Set("languageCode", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("schema_id", suite.SchemaInstanceID)
	rctx.URLParams.Add("module_id", suite.ModuleInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := RemoveModuleFromSchema(req)

	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceSchemaTestSuite) Test00008DeleteSchema() {
	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/schemas", nil)
	req.Header.Set("languageCode", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("schema_id", suite.SchemaInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := DeleteSchema(req)

	deleteModuleToSchema(suite)

	assert.Equal(suite.T(), 200, response.Code)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestServiceSchemaSuite(t *testing.T) {
	suite.Run(t, new(ServiceSchemaTestSuite))
}

func createModuleToSchema(suite *ServiceSchemaTestSuite) {
	data := map[string]interface{}{
		"name":        "Module Teste 01",
		"description": "Descrição do Module Teste 01",
		"code":        "moduleteste01",
		"module":      true,
		"active":      true,
	}
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "http://localhost:3333/api/v1/admin/schemas", bytes.NewBuffer(jsonData))
	req.Header.Set("languageCode", "pt-br")
	req.Header.Set("userID", suite.UserID)

	response := CreateSchema(req)

	schemaValue := reflect.ValueOf(response.Data).Elem()
	suite.ModuleInstanceID = schemaValue.FieldByName("ID").Interface().(string)
}

func deleteModuleToSchema(suite *ServiceSchemaTestSuite) {
	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/schemas", nil)
	req.Header.Set("languageCode", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("schema_id", suite.ModuleInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	DeleteSchema(req)
}
