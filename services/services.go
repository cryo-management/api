package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/andreluzz/go-sql-builder/builder"
	"github.com/andreluzz/go-sql-builder/db"
	"github.com/cryo-management/api/models"
)

//Metadata defines the struct to the api return complementary information to the response data
type Metadata struct {
}

//ResponseError defines the struct to the api response error
type ResponseError struct {
	Code  string `json:"code"`
	Scope string `json:"scope"`
	Error string `json:"erro"`
}

//Response defines the struct to the api response
type Response struct {
	Code     int             `json:"code"`
	Metadata Metadata        `json:"metadata"`
	Data     interface{}     `json:"data"`
	Errors   []ResponseError `json:"errors"`
}

//NewResponse returns an response
func NewResponse() *Response {
	return &Response{
		Code: 200,
	}
}

const (
	//ErrorParsingRequest unable to unmarshall json to struct
	ErrorParsingRequest string = "001-ErrorParsingRequest"
	//ErrorInsertingRecord unable to insert record on database
	ErrorInsertingRecord string = "002-ErrorInsertingRecord"
	//ErrorReturningData unable to return data
	ErrorReturningData string = "003-ErrorReturningData"
	//ErrorDeletingData unable to return data
	ErrorDeletingData string = "004-ErrorDeletingData"
	//ErrorLoadingData unable to load data
	ErrorLoadingData string = "005-ErrorLoadingData"
	//ErrorLogin unable to login user
	ErrorLogin string = "006-ErrorLoginUser"
)

//NewResponseError defines a structure to encode api response data
func NewResponseError(code string, scope, err string) ResponseError {
	return ResponseError{
		Code:  code,
		Scope: scope,
		Error: err,
	}
}

func getColumnsFromBody(body []byte) []string {
	jsonMap := make(map[string]interface{})
	json.Unmarshal(body, &jsonMap)
	columns := []string{}
	for k := range jsonMap {
		columns = append(columns, k)
	}

	return columns
}

func create(r *http.Request, object interface{}, scope, table string) *Response {
	response := NewResponse()
	body, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(body, &object)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorParsingRequest, fmt.Sprintf("%s unmarshal body", scope), err.Error()))

		return response
	}

	id, err := db.InsertStruct(table, object)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorInsertingRecord, fmt.Sprintf("%s create", scope), err.Error()))

		return response
	}

	elementType := reflect.TypeOf(object).Elem()
	elementValue := reflect.ValueOf(object).Elem()
	elementID := elementValue.FieldByName("ID")
	elementID.SetString(id)

	for i := 0; i < elementType.NumField(); i++ {
		if elementType.Field(i).Tag.Get("table") == models.TableTranslations {
			err = models.CreateTranslationsFromStruct(table, r.Header.Get("languageCode"), object)
			if err != nil {
				response.Code = http.StatusInternalServerError
				response.Errors = append(response.Errors, NewResponseError(ErrorInsertingRecord, fmt.Sprintf("%s create translation", scope), err.Error()))

				return response
			}
			break
		}
	}

	response.Data = object

	return response
}

func load(r *http.Request, object interface{}, scope, table string, conditions builder.Builder) *Response {
	response := NewResponse()

	err := db.LoadStruct(table, object, conditions)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorLoadingData, scope, err.Error()))

		return response
	}

	response.Data = object

	return response
}

func remove(r *http.Request, scope, table string, conditions builder.Builder) *Response {
	response := NewResponse()

	err := db.DeleteStruct(table, conditions)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorDeletingData, scope, err.Error()))

		return response
	}

	return response
}
