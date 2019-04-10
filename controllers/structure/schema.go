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
	code := string(chi.URLParam(r, "schema_code"))

	schemaData, errSchema := schema.GetByCode(code)
	if errSchema != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorReturningData, "GetByCode", errSchema.Error()))
		return
	}

	field := new(models.Field)
	errField := field.GetAllWithPermissionBySchemaID(schemaData.ID, &schemaData.Fields)
	if errSchema != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorReturningData, "GetAllWithPermissionBySchemaID", errField.Error()))
		return
	}

	render.JSON(w, r, schemaData)
}
