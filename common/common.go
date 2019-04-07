package common

const (
	//ErrorParsingRequest unable to unmarshall json to struct
	ErrorParsingRequest string = "001-ErrorParsingRequest"
	//ErrorInsertingRecord unable to insert record on database
	ErrorInsertingRecord string = "002-ErrorInsertingRecord"
	//ErrorReturningData unable to return data
	ErrorReturningData string = "003-ErrorReturningData"
)

type responseError struct {
	Code  string `json:"code"`
	Scope string `json:"scope"`
	Error string `json:"erro"`
}

//NewResponseError docs
func NewResponseError(code string, scope, err string) interface{} {
	return responseError{
		Code:  code,
		Scope: scope,
		Error: err,
	}
}
