package schema

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/cryo-management/api/common"
	"github.com/cryo-management/api/structure/translation"

	"github.com/go-chi/render"
)

func postSchema(w http.ResponseWriter, r *http.Request) {
	s := new(Schema)
	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &s)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorParsingRequest, "postSchema", err.Error()))
		return
	}

	id, err := s.Create()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorParsingRequest, "postSchema", err.Error()))
		return
	}

	err = translation.Save(id, r.Header.Get("languageCode"), *s)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorParsingRequest, "postSchema", err.Error()))
		return
	}

	render.JSON(w, r, s)
}

func getAllSchemas(w http.ResponseWriter, r *http.Request) {
	s := new(Schema)

	schemaList, err := s.GetAll()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorReturningData, "getAllSchemas", err.Error()))
		return
	}

	render.JSON(w, r, schemaList)
}
