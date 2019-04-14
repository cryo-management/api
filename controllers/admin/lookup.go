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

func PostLookup(w http.ResponseWriter, r *http.Request) {
	lookup := new(models.Lookup)
	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &lookup)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorParsingRequest, "PostLookup unmarshal body", err.Error()))
		return
	}

	err = lookup.Create()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorInsertingRecord, "PostLookup creating", err.Error()))
		return
	}

	translation := new(models.Translation)
	err = translation.Create(lookup.ID, *lookup)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorInsertingRecord, "PostLookup translation", err.Error()))
		return
	}

	render.JSON(w, r, lookup)
}

func GetLookup(w http.ResponseWriter, r *http.Request) {
	lookup := new(models.Lookup)
	id := string(chi.URLParam(r, "lookup_id"))

	err := lookup.Load(id)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorReturningData, "GetLookup", err.Error()))
		return
	}

	render.JSON(w, r, lookup)
}

func GetAllLookups(w http.ResponseWriter, r *http.Request) {
	lookups := new(models.Lookups)

	err := lookups.Load()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorReturningData, "GetAllLookups load lookups", err.Error()))
		return
	}

	render.JSON(w, r, lookups)
}

func DeleteLookup(w http.ResponseWriter, r *http.Request) {
	lookup := new(models.Lookup)
	id := string(chi.URLParam(r, "lookup_id"))

	err := lookup.Delete(id)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorDeletingData, "DeleteLookup delete lookup", err.Error()))
		return
	}

	translation := new(models.Translation)

	err = translation.DeleteByStructureID(id)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorDeletingData, "DeleteLookup delete translation", err.Error()))
		return
	}

	render.JSON(w, r, id)
}
