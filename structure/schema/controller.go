package schema

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/cryo-management/api/structure/translation"

	"github.com/go-chi/render"
)

func postSchema(w http.ResponseWriter, r *http.Request) {
	s := new(Schema)
	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &s)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, err.Error())
		return
	}

	id, err := s.Create()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, err.Error())
		return
	}

	err = translation.Save(id, r.Header.Get("languageCode"), *s)

	render.JSON(w, r, s)
}
