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

func PostLookup(w http.ResponseWriter, r *http.Request) {
	lookup := new(models.Lookup)
	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &lookup)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorParsingRequest, "PostLookup unmarshal body", err.Error()))
		return
	}

	lookupService := new(services.LookupService)
	err = lookupService.Create(lookup)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorInsertingRecord, "PostLookup creating", err.Error()))
		return
	}

	render.JSON(w, r, lookup)
}

func GetLookup(w http.ResponseWriter, r *http.Request) {
	lookup := new(models.Lookup)
	id := string(chi.URLParam(r, "lookup_id"))

	lookupService := new(services.LookupService)
	err := lookupService.Load(lookup, id)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorReturningData, "GetLookup load", err.Error()))
		return
	}

	render.JSON(w, r, lookup)
}

func GetAllLookups(w http.ResponseWriter, r *http.Request) {
	lookups := new(models.Lookups)

	lookupService := new(services.LookupService)
	err := lookupService.LoadAll(lookups)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorReturningData, "GetAllLookups load", err.Error()))
		return
	}

	render.JSON(w, r, lookups)
}

func DeleteLookup(w http.ResponseWriter, r *http.Request) {
	lookup := new(models.Lookup)
	id := string(chi.URLParam(r, "lookup_id"))

	lookupService := new(services.LookupService)
	err := lookupService.Delete(lookup, id)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorDeletingData, "DeleteLookup delete lookup", err.Error()))
		return
	}

	render.JSON(w, r, id)
}
