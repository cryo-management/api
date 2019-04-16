package services

type Metadata struct {
}

type ResponseError struct {
	Code  string `json:"code"`
	Scope string `json:"scope"`
	Error string `json:"erro"`
}

type Response struct {
	Metadata Metadata        `json:"metadata"`
	Data     []interface{}   `json:"data"`
	Errors   []ResponseError `json:"errors"`
}
