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

func PostField(w http.ResponseWriter, r *http.Request) {
	field := new(models.Field)
	schemaID := string(chi.URLParam(r, "schema_id"))
	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &field)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorParsingRequest, "PostField unmarshal body", err.Error()))
		return
	}

	field.SchemaID = schemaID

	err = field.Create()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorInsertingRecord, "PostField creating", err.Error()))
		return
	}

	translation := new(models.Translation)
	err = translation.Create(field.ID, *field)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorInsertingRecord, "PostField translation", err.Error()))
		return
	}

	render.JSON(w, r, field)
}

func GetField(w http.ResponseWriter, r *http.Request) {
	field := new(models.Field)
	id := string(chi.URLParam(r, "field_id"))

	err := field.Load(id)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorReturningData, "GetField", err.Error()))
		return
	}

	render.JSON(w, r, field)
}

func GetAllFields(w http.ResponseWriter, r *http.Request) {
	fields := new(models.Fields)
	schemaID := string(chi.URLParam(r, "schema_id"))

	err := fields.Load(schemaID)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorReturningData, "GetAllFields load fields", err.Error()))
		return
	}

	render.JSON(w, r, fields)
}

func DeleteField(w http.ResponseWriter, r *http.Request) {
	field := new(models.Field)
	id := string(chi.URLParam(r, "field_id"))

	err := field.Delete(id)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorDeletingData, "DeleteField delete field", err.Error()))
		return
	}

	translation := new(models.Translation)

	err = translation.DeleteByStructureID(id)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorDeletingData, "DeleteField delete translation", err.Error()))
		return
	}

	render.JSON(w, r, id)
}