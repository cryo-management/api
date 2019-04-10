package admin

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/cryo-management/api/common"
	"github.com/cryo-management/api/models"

	"github.com/go-chi/render"
)

//GetSchema docs
func GetSchema(w http.ResponseWriter, r *http.Request) {
	schema := new(models.Schema)
	code := string(chi.URLParam(r, "schema_code"))

	schemaData, err := schema.GetByCode(code)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorReturningData, "GetAllSchemas", err.Error()))
		return
	}

	render.JSON(w, r, schemaData)
}

//PostSchema docs
func PostSchema(w http.ResponseWriter, r *http.Request) {
	schema := new(models.Schema)
	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &schema)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorParsingRequest, "PostSchema unmarshal body", err.Error()))
		return
	}

	id, err := schema.Create()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorInsertingRecord, "PostSchema creating", err.Error()))
		return
	}

	translation := new(models.Translation)
	err = translation.Create(id, r.Header.Get("languageCode"), *schema)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorInsertingRecord, "PostSchema translation", err.Error()))
		return
	}

	render.JSON(w, r, schema)
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
