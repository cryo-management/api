package sctructure

import (
	"net/http"

	"github.com/cryo-management/api/common"
	"github.com/cryo-management/api/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

//GetSchema docs
func GetSchema(w http.ResponseWriter, r *http.Request) {
	schema := new(models.Schema)
	id := string(chi.URLParam(r, "schema_id"))

	err := schema.Load(id)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorReturningData, "GetSchema load schema", err.Error()))
		return
	}

	err = schema.Fields.LoadByPermission(schema.ID)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorReturningData, "GetSchema load fields", err.Error()))
		return
	}

	render.JSON(w, r, schema)
}
