package sctructure

import "net/http"

//GetSchema docs
func GetSchema(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get schema structure instance"))
}
