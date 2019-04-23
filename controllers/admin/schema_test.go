package admin

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andreluzz/go-sql-builder/db"
	"github.com/cryo-management/api/config"
	"github.com/cryo-management/api/models"
	"github.com/cryo-management/api/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ControllerSchemaTestSuite struct {
	suite.Suite
	InstanceID string
}

func (suite *ControllerSchemaTestSuite) SetupTest() {
	config, _ := config.NewConfig("..\\..\\config.toml")
	db.Connect(config.Host, config.Port, config.User, config.Password, config.DBName, false)
}

func (suite *ControllerSchemaTestSuite) Test00001CreateSchema() {
	data := map[string]interface{}{
		"name":        "Schema Teste 01",
		"description": "Descrição do Schema Teste 01",
		"code":        "schemateste01",
		"module":      false,
		"active":      true,
	}
	jsonData, _ := json.Marshal(data)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "http://localhost:3333/api/v1/admin/schemas", bytes.NewBuffer(jsonData))
	req.Header.Set("LanguageCode", "pt-br")

	PostSchema(res, req)

	body, _ := ioutil.ReadAll(res.Body)
	response := services.Response{}
	json.Unmarshal(body, &response)

	schema := models.Schema{}
	jsonMap, _ := json.Marshal(response.Data)
	json.Unmarshal(jsonMap, &schema)

	assert.NotNil(suite.T(), response.Data, "response.Data should not be null")

	assert.Equal(suite.T(), 200, response.Code)

	suite.InstanceID = schema.ID
}

// func (suite *ControllerSchemaTestSuite) Test00002LoadAllSchemas() {
// 	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/schemas", nil)
// 	req.Header.Set("LanguageCode", "pt-br")

// 	response := LoadAllSchemas(req)

// 	result := response.Data != nil && response.Code == 200

// 	assert.Equal(suite.T(), result, true)
// }

// func (suite *ControllerSchemaTestSuite) Test00003LoadSchema() {
// 	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/schemas", nil)
// 	req.Header.Set("LanguageCode", "pt-br")

// 	rctx := chi.NewRouteContext()
// 	rctx.URLParams.Add("schema_id", suite.InstanceID)
// 	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

// 	response := LoadSchema(req)

// 	result := response.Data != nil && response.Code == 200

// 	assert.Equal(suite.T(), result, true)
// }

// func (suite *ControllerSchemaTestSuite) Test00004UpdateSchema() {
// 	data := map[string]interface{}{
// 		"description": "Descrição do Schema Teste 01 Updated",
// 		"active":      false,
// 	}
// 	jsonData, _ := json.Marshal(&data)

// 	req, _ := http.NewRequest("PATCH", "http://localhost:3333/api/v1/admin/schemas", bytes.NewBuffer(jsonData))
// 	req.Header.Set("LanguageCode", "pt-br")

// 	rctx := chi.NewRouteContext()
// 	rctx.URLParams.Add("schema_id", suite.InstanceID)
// 	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

// 	response := UpdateSchema(req)

// 	result := response.Code == 200

// 	assert.Equal(suite.T(), result, true)
// }

// func (suite *ControllerSchemaTestSuite) Test00005DeleteSchema() {
// 	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/schemas", nil)
// 	req.Header.Set("LanguageCode", "pt-br")

// 	rctx := chi.NewRouteContext()
// 	rctx.URLParams.Add("schema_id", suite.InstanceID)
// 	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

// 	response := DeleteSchema(req)

// 	result := response.Code == 200

// 	assert.Equal(suite.T(), result, true)
// }

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestControllerSchemaSuite(t *testing.T) {
	suite.Run(t, new(ControllerSchemaTestSuite))
}
