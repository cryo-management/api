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

func PostLookupOption(w http.ResponseWriter, r *http.Request) {
	lookupOption := new(models.LookupOption)
	lookupID := string(chi.URLParam(r, "lookup_id"))
	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &lookupOption)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorParsingRequest, "PostLookupOption unmarshal body", err.Error()))
		return
	}

	lookupOption.LookupID = lookupID

	err = lookupOption.Create()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorInsertingRecord, "PostLookupOption creating", err.Error()))
		return
	}

	translation := new(models.Translation)
	err = translation.Create(lookupOption.ID, *lookupOption)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorInsertingRecord, "PostLookupOption translation", err.Error()))
		return
	}

	render.JSON(w, r, lookupOption)
}

func GetLookupOption(w http.ResponseWriter, r *http.Request) {
	lookupOption := new(models.LookupOption)
	id := string(chi.URLParam(r, "lookup_option_id"))

	err := lookupOption.Load(id)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorReturningData, "GetLookupOption", err.Error()))
		return
	}

	render.JSON(w, r, lookupOption)
}

func GetAllLookupOptions(w http.ResponseWriter, r *http.Request) {
	lookupOptions := new(models.LookupOptions)
	lookupID := string(chi.URLParam(r, "lookup_id"))

	err := lookupOptions.Load(lookupID)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorReturningData, "GetAllLookupOptions load lookupOptions", err.Error()))
		return
	}

	render.JSON(w, r, lookupOptions)
}

func DeleteLookupOption(w http.ResponseWriter, r *http.Request) {
	lookupOption := new(models.LookupOption)
	id := string(chi.URLParam(r, "lookup_option_id"))

	err := lookupOption.Delete(id)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorDeletingData, "DeleteLookupOption delete lookupOption", err.Error()))
		return
	}

	translation := new(models.Translation)

	err = translation.DeleteByStructureID(id)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorDeletingData, "DeleteLookupOption delete translation", err.Error()))
		return
	}

	render.JSON(w, r, id)
}
