package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/cryo-management/api/models"
)

func CreateSchema(r *http.Request) *Response {
	response := new(Response)
	body, _ := ioutil.ReadAll(r.Body)
	schema := new(models.Schema)
	err := json.Unmarshal(body, &schema)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorParsingRequest, "CreateSchema unmarshal body", err.Error()))
		return response
	}

	err = schema.Create()
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorParsingRequest, "CreateSchema create", err.Error()))
		return response
	}

	// translation := new(models.Translation)
	// err = translation.Create(schema)
	// if err != nil {
	// 	response.Code = http.StatusInternalServerError
	// 	response.Errors = append(response.Errors, NewResponseError(ErrorParsingRequest, "CreateSchema create translation", err.Error()))
	// 	return response
	// }

	response.Code = 200
	response.Data = schema

	return response
}

// func (s *SchemaService) Load(schema *models.Schema, id string) error {
// 	err := schema.Load(id)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (s *SchemaService) LoadAll(schemas *models.Schemas) error {
// 	err := schemas.Load()
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (s *SchemaService) Delete(schema *models.Schema, id string) error {
// 	err := schema.Delete(id)
// 	if err != nil {
// 		return err
// 	}

// 	translationService := new(TranslationService)
// 	err = translationService.DeleteByStructureID(id)
// 	if err != nil {
// 		return err
// 	}

// 	groupPermission := new(models.GroupPermission)
// 	groupPermission.StructureID = id

// 	groupService := new(GroupService)
// 	err = groupService.DeletePermission(groupPermission)
// 	if err != nil {
// 		return err
// 	}

// 	fields := new(models.Fields)

// 	fieldService := new(FieldService)
// 	err = fieldService.DeleteBySchemaID(fields, id)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
