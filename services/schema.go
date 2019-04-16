package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/cryo-management/api/models"
)

//CreateSchema persists the request body creating a new schema in the database
func CreateSchema(r *http.Request) *Response {
	response := NewResponse()
	body, _ := ioutil.ReadAll(r.Body)
	schema := new(models.Schema)
	err := json.Unmarshal(body, &schema)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorParsingRequest, "CreateSchema unmarshal body", err.Error()))
		return response
	}

	id, err := schema.Create()
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorInsertingRecord, "CreateSchema create", err.Error()))
		return response
	}
	schema.ID = id

	err = models.CreateTranslationsFromStruct(models.TableSchema, "pt-br", schema)

	response.Data = schema

	return response
}
