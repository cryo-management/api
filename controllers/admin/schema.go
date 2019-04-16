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

func PostSchema(w http.ResponseWriter, r *http.Request) {
	schema := new(models.Schema)
	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &schema)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorParsingRequest, "PostSchema unmarshal body", err.Error()))
		return
	}

	schemaService := new(services.SchemaService)
	err = schemaService.Create(schema)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorInsertingRecord, "PostSchema creating", err.Error()))
		return
	}

	render.JSON(w, r, schema)
}

func GetSchema(w http.ResponseWriter, r *http.Request) {
	schema := new(models.Schema)
	id := string(chi.URLParam(r, "schema_id"))

	schemaService := new(services.SchemaService)
	err := schemaService.Load(schema, id)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorReturningData, "GetSchema load", err.Error()))
		return
	}

	render.JSON(w, r, schema)
}

func GetAllSchemas(w http.ResponseWriter, r *http.Request) {
	schemas := new(models.Schemas)

	schemaService := new(services.SchemaService)
	err := schemaService.LoadAll(schemas)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorReturningData, "GetAllSchemas load", err.Error()))
		return
	}

	render.JSON(w, r, schemas)
}

func DeleteSchema(w http.ResponseWriter, r *http.Request) {
	schema := new(models.Schema)
	id := string(chi.URLParam(r, "schema_id"))

	schemaService := new(services.SchemaService)
	err := schemaService.Delete(schema, id)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorDeletingData, "DeleteSchema delete schema", err.Error()))
		return
	}

	render.JSON(w, r, id)
}
