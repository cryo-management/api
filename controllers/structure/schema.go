package sctructure

// func GetSchema(w http.ResponseWriter, r *http.Request) {
// 	schema := new(models.Schema)
// 	id := string(chi.URLParam(r, "schema_id"))

// 	schemaService := new(services.SchemaService)
// 	err := schemaService.Load(schema, id)
// 	if err != nil {
// 		render.Status(r, http.StatusInternalServerError)
// 		render.JSON(w, r, common.NewResponseError(common.ErrorReturningData, "GetSchema load schema", err.Error()))
// 		return
// 	}

// 	render.JSON(w, r, schema)
// }
