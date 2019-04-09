package sctructure

import "net/http"

//GetSchema docs
func GetSchema(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get schema structure instance"))
	//get schema by code
	//s.GetSchema("schema_code")
	//get fields by schema id
	//err = f.GetFields(s.ID, *s.Fields)
	//render(w, r, s)
}
