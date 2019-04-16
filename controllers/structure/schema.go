package sctructure

import (
	"net/http"

	"github.com/cryo-management/api/common"
	"github.com/cryo-management/api/models"
	services "github.com/cryo-management/api/services/structure"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func GetSchema(w http.ResponseWriter, r *http.Request) {
	schema := new(models.Schema)
	id := string(chi.URLParam(r, "schema_id"))

	schemaService := new(services.SchemaService)
	err := schemaService.Load(schema, id)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorReturningData, "GetSchema load schema", err.Error()))
		return
	}

	render.JSON(w, r, schema)
}
