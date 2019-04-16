package admin

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/cryo-management/api/common"
	"github.com/cryo-management/api/models"
	services "github.com/cryo-management/api/services/admin"

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

	fieldService := new(services.FieldService)
	err = fieldService.Create(field)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorInsertingRecord, "PostField creating", err.Error()))
		return
	}

	render.JSON(w, r, field)
}

func GetField(w http.ResponseWriter, r *http.Request) {
	field := new(models.Field)
	id := string(chi.URLParam(r, "field_id"))

	fieldService := new(services.FieldService)
	err := fieldService.Load(field, id)
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

	fieldService := new(services.FieldService)
	err := fieldService.LoadAll(fields, schemaID)
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

	fieldService := new(services.FieldService)
	err := fieldService.Delete(field, id)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorDeletingData, "DeleteSchema delete schema", err.Error()))
		return
	}

	render.JSON(w, r, id)
}
