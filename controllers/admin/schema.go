package admin

import (
	"net/http"

	"github.com/cryo-management/api/services"

	"github.com/go-chi/render"
)

func PostSchema(w http.ResponseWriter, r *http.Request) {
	response := services.CreateSchema(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// func GetSchema(w http.ResponseWriter, r *http.Request) {
// 	schema := new(models.Schema)
// 	id := string(chi.URLParam(r, "schema_id"))

// 	schemaService := new(services.SchemaService)
// 	err := schemaService.Load(schema, id)
// 	if err != nil {
// 		render.Status(r, http.StatusInternalServerError)
// 		render.JSON(w, r, common.NewResponseError(common.ErrorReturningData, "GetSchema load", err.Error()))
// 		return
// 	}

// 	render.JSON(w, r, schema)
// }

// func GetAllSchemas(w http.ResponseWriter, r *http.Request) {
// 	schemas := new(models.Schemas)

// 	schemaService := new(services.SchemaService)
// 	err := schemaService.LoadAll(schemas)
// 	if err != nil {
// 		render.Status(r, http.StatusInternalServerError)
// 		render.JSON(w, r, common.NewResponseError(common.ErrorReturningData, "GetAllSchemas load", err.Error()))
// 		return
// 	}

// 	render.JSON(w, r, schemas)
// }

// func DeleteSchema(w http.ResponseWriter, r *http.Request) {
// 	schema := new(models.Schema)
// 	id := string(chi.URLParam(r, "schema_id"))

// 	schemaService := new(services.SchemaService)
// 	err := schemaService.Delete(schema, id)
// 	if err != nil {
// 		render.Status(r, http.StatusInternalServerError)
// 		render.JSON(w, r, common.NewResponseError(common.ErrorDeletingData, "DeleteSchema delete schema", err.Error()))
// 		return
// 	}

// 	render.JSON(w, r, id)
// }
