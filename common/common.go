package common

const (
	//ErrorParsingRequest unable to unmarshall json to struct
	ErrorParsingRequest string = "001-ErrorParsingRequest"
	//ErrorInsertingRecord unable to insert record on database
	ErrorInsertingRecord string = "002-ErrorInsertingRecord"
	//ErrorReturningData unable to return data
	ErrorReturningData string = "003-ErrorReturningData"
	//ErrorDeletingData unable to return data
	ErrorDeletingData string = "004-ErrorDeletingData"
)

type responseError struct {
	Code  string `json:"code"`
	Scope string `json:"scope"`
	Error string `json:"erro"`
}

func NewResponseError(code string, scope, err string) interface{} {
	return responseError{
		Code:  code,
		Scope: scope,
		Error: err,
	}
}
