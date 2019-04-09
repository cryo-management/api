package admin

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/cryo-management/api/common"
	"github.com/cryo-management/api/models"

	"github.com/go-chi/render"
)

//GetSchema docs
func GetSchema(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get schema structure instance"))
}

//PostSchema docs
func PostSchema(w http.ResponseWriter, r *http.Request) {
	s := new(models.Schema)
	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &s)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorParsingRequest, "PostSchema unmarshal body", err.Error()))
		return
	}

	id, err := s.Create()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorInsertingRecord, "PostSchema creating", err.Error()))
		return
	}

	err = models.CreateTranslation(id, r.Header.Get("languageCode"), *s)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorInsertingRecord, "PostSchema translation", err.Error()))
		return
	}

	render.JSON(w, r, s)
}

//GetAllSchemas docs
func GetAllSchemas(w http.ResponseWriter, r *http.Request) {
	s := new(models.Schema)

	schemaList, err := s.GetAll()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorReturningData, "GetAllSchemas", err.Error()))
		return
	}

	render.JSON(w, r, schemaList)
}
