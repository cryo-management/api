package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"time"

	"github.com/andreluzz/go-sql-builder/builder"
	"github.com/andreluzz/go-sql-builder/db"
	"github.com/cryo-management/api/models"
)

// Metadata defines the struct to the api return complementary information to the response data
type Metadata struct {
}

// ResponseError defines the struct to the api response error
type ResponseError struct {
	Code  string `json:"code"`
	Scope string `json:"scope"`
	Error string `json:"erro"`
}

// Response defines the struct to the api response
type Response struct {
	Code     int             `json:"code"`
	Metadata Metadata        `json:"metadata"`
	Data     interface{}     `json:"data"`
	Errors   []ResponseError `json:"errors"`
}

// NewResponse returns an response
func NewResponse() *Response {
	return &Response{
		Code: 200,
	}
}

const (
	// ErrorParsingRequest unable to unmarshall json to struct
	ErrorParsingRequest string = "001-ErrorParsingRequest"
	// ErrorInsertingRecord unable to insert record on database
	ErrorInsertingRecord string = "002-ErrorInsertingRecord"
	// ErrorReturningData unable to return data
	ErrorReturningData string = "003-ErrorReturningData"
	// ErrorDeletingData unable to return data
	ErrorDeletingData string = "004-ErrorDeletingData"
	// ErrorLoadingData unable to load data
	ErrorLoadingData string = "005-ErrorLoadingData"
	// ErrorLogin unable to login user
	ErrorLogin string = "006-ErrorLoginUser"
)

// NewResponseError defines a structure to encode api response data
func NewResponseError(code string, scope, err string) ResponseError {
	return ResponseError{
		Code:  code,
		Scope: scope,
		Error: err,
	}
}

// GetUpdateColumnsFromBody get request body and return an string array with columns from the body
func GetUpdateColumnsFromBody(body []byte) []string {
	jsonMap := make(map[string]interface{})
	json.Unmarshal(body, &jsonMap)
	columns := []string{}
	for k := range jsonMap {
		if k != "created_by" && k != "created_at" && k != "updated_by" && k != "updated_at" {
			columns = append(columns, k)
		}
	}
	columns = append(columns, "updated_by")
	columns = append(columns, "updated_at")

	return columns
}

// GetFilterColumns return translation columns from the object
func GetFilterColumns(r *http.Request, object interface{}, table string) (map[string]interface{}, error) {
	query := r.URL.Query()
	jsonFilters := query.Get("filter")
	filterColumns := make(map[string]interface{})

	if jsonFilters != "" {
		data := []byte(jsonFilters)
		filterMap := make(map[string]interface{})
		err := json.Unmarshal(data, &filterMap)

		if err != nil {
			return nil, err
		}

		elementType := reflect.TypeOf(object).Elem()

		if elementType.Kind() == reflect.Slice {
			elementType = elementType.Elem()
		}

		for filter, value := range filterMap {
			column := fmt.Sprintf("%s.%s", table, filter)
			for i := 0; i < elementType.NumField(); i++ {
				elementField := elementType.Field(i)
				if filter == elementField.Tag.Get("json") && elementField.Tag.Get("table") == models.TableCoreTranslations {
					column = fmt.Sprintf("%s.%s", elementField.Tag.Get("alias"), "value")
					break
				}
			}
			filterColumns[column] = value
		}
	}

	return filterColumns, nil
}

// GetTranslationLanguageCodeColumns return translation columns from the object
func GetTranslationLanguageCodeColumns(object interface{}, columns ...string) []string {
	translationColumns := []string{}
	elementType := reflect.TypeOf(object).Elem()

	if elementType.Kind() == reflect.Slice {
		elementType = elementType.Elem()
	}

	for i := 0; i < elementType.NumField(); i++ {
		elementField := elementType.Field(i)
		if elementField.Tag.Get("table") == models.TableCoreTranslations {
			jsonColumn := elementField.Tag.Get("json")
			translationTableAlias := elementField.Tag.Get("alias")
			if len(columns) > 0 {
				for _, column := range columns {
					if column == jsonColumn {
						translationColumns = append(translationColumns, fmt.Sprintf("%s.language_code", translationTableAlias))
					}
				}
			} else {
				translationColumns = append(translationColumns, fmt.Sprintf("%s.language_code", translationTableAlias))
			}
		}
	}

	return translationColumns
}

// Create object data in the database
func Create(r *http.Request, object interface{}, scope, table string) *Response {
	response := NewResponse()
	body, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(body, &object)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorParsingRequest, fmt.Sprintf("%s unmarshal body", scope), err.Error()))

		return response
	}

	userID := r.Header.Get("userID")
	now := time.Now()
	elementValue := reflect.ValueOf(object).Elem()
	elementCreatedBy := elementValue.FieldByName("CreatedBy")
	elementUpdatedBy := elementValue.FieldByName("UpdatedBy")
	elementCreatedAt := elementValue.FieldByName("CreatedAt")
	elementUpdatedAt := elementValue.FieldByName("UpdatedAt")
	elementCreatedBy.SetString(userID)
	elementUpdatedBy.SetString(userID)
	elementCreatedAt.Set(reflect.ValueOf(now))
	elementUpdatedAt.Set(reflect.ValueOf(now))

	id, err := db.InsertStruct(table, object)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorInsertingRecord, fmt.Sprintf("%s create", scope), err.Error()))

		return response
	}

	elementID := elementValue.FieldByName("ID")
	elementID.SetString(id)

	translationColumns := GetTranslationLanguageCodeColumns(object)

	if len(translationColumns) > 0 {
		err = models.CreateTranslationsFromStruct(table, r.Header.Get("userID"), r.Header.Get("Content-Language"), object)
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Errors = append(response.Errors, NewResponseError(ErrorInsertingRecord, fmt.Sprintf("%s create translation", scope), err.Error()))

			return response
		}
	}

	response.Data = object

	return response
}

// Load object data from the database
func Load(r *http.Request, object interface{}, scope, table string, conditions builder.Builder) *Response {
	response := NewResponse()

	filterColumns, _ := GetFilterColumns(r, object, table)

	if len(filterColumns) > 0 {
		newCondition := []builder.Builder{}
		if conditions != nil {
			newCondition = append(newCondition, conditions)
		}
		for column, value := range filterColumns {
			newCondition = append(newCondition, builder.Equal(column, value))
		}
		conditions = builder.And(newCondition...)
	}

	translationColumns := GetTranslationLanguageCodeColumns(object)

	if len(translationColumns) > 0 {
		newCondition := []builder.Builder{}
		if conditions != nil {
			newCondition = append(newCondition, conditions)
		}
		for _, translationColumn := range translationColumns {
			newCondition = append(newCondition, builder.Equal(translationColumn, r.Header.Get("Content-Language")))
		}
		conditions = builder.And(newCondition...)
	}

	err := db.LoadStruct(table, object, conditions)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorLoadingData, scope, err.Error()))

		return response
	}

	response.Data = object

	return response
}

// Remove object data from the database
func Remove(r *http.Request, scope, table string, conditions builder.Builder) *Response {
	response := NewResponse()

	err := db.DeleteStruct(table, conditions)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorDeletingData, scope, err.Error()))

		return response
	}

	return response
}

// Update object data in the database
func Update(r *http.Request, object interface{}, scope, table string, condition builder.Builder) *Response {
	response := NewResponse()
	body, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(body, &object)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorParsingRequest, fmt.Sprintf("%s unmarshal body", scope), err.Error()))

		return response
	}

	columns := GetUpdateColumnsFromBody(body)

	userID := r.Header.Get("userID")
	now := time.Now()
	elementValue := reflect.ValueOf(object).Elem()
	elementUpdatedBy := elementValue.FieldByName("UpdatedBy")
	elementUpdatedAt := elementValue.FieldByName("UpdatedAt")
	elementUpdatedBy.SetString(userID)
	elementUpdatedAt.Set(reflect.ValueOf(now))

	err = db.UpdateStruct(table, object, condition, columns...)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorInsertingRecord, fmt.Sprintf("%s update", scope), err.Error()))

		return response
	}

	translationColumns := GetTranslationLanguageCodeColumns(object, columns...)

	if len(translationColumns) > 0 {
		err = models.UpdateTranslationsFromStruct(table, r.Header.Get("userID"), r.Header.Get("Content-Language"), object, columns...)
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Errors = append(response.Errors, NewResponseError(ErrorInsertingRecord, fmt.Sprintf("%s update translation", scope), err.Error()))

			return response
		}
	}

	response.Data = object

	return response
}
