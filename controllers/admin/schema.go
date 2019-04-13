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

	err = schema.Create()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorInsertingRecord, "PostSchema creating", err.Error()))
		return
	}

	translation := new(models.Translation)
	err = translation.Create(schema.ID, *schema)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorInsertingRecord, "PostSchema translation", err.Error()))
		return
	}

	render.JSON(w, r, schema)
}

//GetSchema docs
func GetSchema(w http.ResponseWriter, r *http.Request) {
	schema := new(models.Schema)
	code := string(chi.URLParam(r, "schema_code"))

	err := schema.Load(code)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorReturningData, "GetSchema load", err.Error()))
		return
	}

	render.JSON(w, r, schema)
}

//GetAllSchemas docs
func GetAllSchemas(w http.ResponseWriter, r *http.Request) {
	schemas := new(models.Schemas)

	err := schemas.Load()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorReturningData, "GetAllSchemas load", err.Error()))
		return
	}

	render.JSON(w, r, schemas)
}

//DeleteSchema docs
func DeleteSchema(w http.ResponseWriter, r *http.Request) {
	schema := new(models.Schema)
	id := string(chi.URLParam(r, "schema_id"))

	err := schema.Delete(id)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorDeletingData, "DeleteSchema delete schema", err.Error()))
		return
	}

	translation := new(models.Translation)

	err = translation.DeleteByStructureID(id)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorDeletingData, "DeleteSchema delete translation", err.Error()))
		return
	}

	render.JSON(w, r, id)
}
