package services

import (
	"net/http"

	"github.com/andreluzz/go-sql-builder/builder"
	"github.com/andreluzz/go-sql-builder/db"
)

type Metadata struct {
}

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
)

//NewResponseError defines a structure to encode api response data
func NewResponseError(code string, scope, err string) ResponseError {
	return ResponseError{
		Code:  code,
		Scope: scope,
		Error: err,
	}
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
