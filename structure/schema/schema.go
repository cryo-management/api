package schema

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cryo-management/api/db"
	"github.com/go-chi/render"
)

//Schema docs
type Schema struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Code        string `json:"code"`
	Module      bool   `json:"module"`
	Active      bool   `json:"active"`
}

func postSchema(w http.ResponseWriter, r *http.Request) {

	/*
		TODO create validation. Check if is better to use middleware
		err := validate.Run(r.Header.Get("userId"), validate.CreateSchema)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, err.Error())
			return
		}
	*/

	var s Schema
	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &s)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, err.Error())
		return
	}

	db := new(db.Database)
	id, err := db.Insert("insert into schemas (code, module, active) values ($1, $2, $3)", s.Code, s.Module, s.Active)
	fmt.Println(s)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, err.Error())
		return
	}
	s.ID = id

	query := "insert into translations (structure_type, structure_id, structure_code, value, language_code) values ($1, $2, $3, $4, $5)"

	id, err = db.Insert(query, "schemas", s.ID, "name", s.Name, r.Header.Get("languageCode"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, err.Error())
		return
	}

	id, err = db.Insert(query, "schemas", s.ID, "description", s.Description, r.Header.Get("languageCode"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, err.Error())
		return
	}

	render.JSON(w, r, s)
}
