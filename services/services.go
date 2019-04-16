package services

type Metadata struct {
}

type ResponseError struct {
	Code  string `json:"code"`
	Scope string `json:"scope"`
	Error string `json:"erro"`
}

type Response struct {
	Code     int             `json:"code"`
	Metadata Metadata        `json:"metadata"`
	Data     interface{}     `json:"data"`
	Errors   []ResponseError `json:"errors"`
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
)

func NewResponseError(code string, scope, err string) ResponseError {
	return ResponseError{
		Code:  code,
		Scope: scope,
		Error: err,
	}
}
