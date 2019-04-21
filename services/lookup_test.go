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

type ServiceLookupTestSuite struct {
	suite.Suite
	LookupInstanceID       string
	LookupOptionInstanceID string
}

func (suite *ServiceLookupTestSuite) SetupTest() {
	pathSeparator := string(os.PathSeparator)
	file := fmt.Sprintf("..%sconfig.toml", pathSeparator)
	config := config.Config{}
	toml.DecodeFile(file, &config)
	db.Connect(config.Host, config.Port, config.User, config.Password, config.DBName, false)
}

func (suite *ServiceLookupTestSuite) Test00001CreateLookup() {
	data := map[string]interface{}{
		"name":         "Lookup Teste 01",
		"description":  "Descrição do Lookup Teste 01",
		"code":         "lookupteste01",
		"type":         "static",
		"value":        "id",
		"label":        "value",
		"autocomplete": "value",
		"active":       true,
	}
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "http://localhost:3333/api/v1/admin/lookups", bytes.NewBuffer(jsonData))
	req.Header.Set("LanguageCode", "pt-br")

	response := CreateLookup(req)

	result := response.Data != nil && response.Code == 200
	lookupValue := reflect.ValueOf(response.Data).Elem()
	suite.LookupInstanceID = lookupValue.FieldByName("ID").Interface().(string)

	assert.Equal(suite.T(), result, true)
}

func (suite *ServiceLookupTestSuite) Test00002LoadAllLookups() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/lookups", nil)
	req.Header.Set("LanguageCode", "pt-br")

	response := LoadAllLookups(req)

	result := response.Data != nil && response.Code == 200

	assert.Equal(suite.T(), result, true)
}

func (suite *ServiceLookupTestSuite) Test00003LoadLookup() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/lookups", nil)
	req.Header.Set("LanguageCode", "pt-br")

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("lookup_id", suite.LookupInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := LoadLookup(req)

	result := response.Data != nil && response.Code == 200

	assert.Equal(suite.T(), result, true)
}

func (suite *ServiceLookupTestSuite) Test00004UpdateLookup() {
	data := map[string]interface{}{
		"description": "Descrição do Lookup Teste 01 Updated",
		"active":      false,
	}
	jsonData, _ := json.Marshal(&data)

	req, _ := http.NewRequest("PATCH", "http://localhost:3333/api/v1/admin/lookups", bytes.NewBuffer(jsonData))
	req.Header.Set("LanguageCode", "pt-br")

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("lookup_id", suite.LookupInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := UpdateLookup(req)

	result := response.Code == 200

	assert.Equal(suite.T(), result, true)
}

func (suite *ServiceLookupTestSuite) Test00005CreateLookupOption() {
	data := map[string]interface{}{
		"lookup_id": suite.LookupInstanceID,
		"value":     "valorteste01",
		"label":     "Valor Teste 01",
		"active":    true,
	}
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "http://localhost:3333/api/v1/admin/lookups", bytes.NewBuffer(jsonData))
	req.Header.Set("LanguageCode", "pt-br")

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("lookup_id", suite.LookupInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := CreateLookupOption(req)

	result := response.Data != nil && response.Code == 200
	lookupOptionValue := reflect.ValueOf(response.Data).Elem()
	suite.LookupOptionInstanceID = lookupOptionValue.FieldByName("ID").Interface().(string)

	assert.Equal(suite.T(), result, true)
}

func (suite *ServiceLookupTestSuite) Test00006LoadAllLookupOptions() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/lookups", nil)
	req.Header.Set("LanguageCode", "pt-br")

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("lookup_id", suite.LookupInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := LoadAllLookupOptions(req)

	result := response.Data != nil && response.Code == 200

	assert.Equal(suite.T(), result, true)
}

func (suite *ServiceLookupTestSuite) Test00007LoadLookupOption() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/lookups", nil)
	req.Header.Set("LanguageCode", "pt-br")

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("lookup_id", suite.LookupInstanceID)
	rctx.URLParams.Add("lookup_option_id", suite.LookupOptionInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := LoadLookupOption(req)

	result := response.Data != nil && response.Code == 200

	assert.Equal(suite.T(), result, true)
}

func (suite *ServiceLookupTestSuite) Test00008UpdateLookupOption() {
	data := map[string]interface{}{
		"label":  "Valor Teste 01 Updated",
		"active": true,
	}
	jsonData, _ := json.Marshal(&data)

	req, _ := http.NewRequest("PATCH", "http://localhost:3333/api/v1/admin/lookups", bytes.NewBuffer(jsonData))
	req.Header.Set("LanguageCode", "pt-br")

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("lookup_id", suite.LookupInstanceID)
	rctx.URLParams.Add("lookup_option_id", suite.LookupOptionInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := UpdateLookupOption(req)

	result := response.Code == 200

	assert.Equal(suite.T(), result, true)
}

func (suite *ServiceLookupTestSuite) Test00009DeleteLookupOption() {
	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/lookups", nil)
	req.Header.Set("LanguageCode", "pt-br")

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("lookup_id", suite.LookupInstanceID)
	rctx.URLParams.Add("lookup_option_id", suite.LookupOptionInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := DeleteLookupOption(req)

	result := response.Code == 200

	assert.Equal(suite.T(), result, true)
}

func (suite *ServiceLookupTestSuite) Test00010DeleteLookup() {
	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/lookups", nil)
	req.Header.Set("LanguageCode", "pt-br")

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("lookup_id", suite.LookupInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := DeleteLookup(req)

	result := response.Code == 200

	assert.Equal(suite.T(), result, true)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestServiceLookupSuite(t *testing.T) {
	suite.Run(t, new(ServiceLookupTestSuite))
}
