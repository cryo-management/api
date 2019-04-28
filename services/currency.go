package services

import (
	"fmt"
	"net/http"

	"github.com/andreluzz/go-sql-builder/builder"
	"github.com/go-chi/chi"

	"github.com/cryo-management/api/models"
)

// CreateCurrency persists the request body creating a new object in the database
func CreateCurrency(r *http.Request) *Response {
	currency := models.Currency{}

	return create(r, &currency, "CreateCurrency", models.TableCoreCurrencies)
}

// LoadAllCurrencies return all instances from the object
func LoadAllCurrencies(r *http.Request) *Response {
	currencies := []models.Currency{}

	return load(r, &currencies, "LoadAllCurrencies", models.TableCoreCurrencies, nil)
}

// LoadCurrency return only one object from the database
func LoadCurrency(r *http.Request) *Response {
	currency := models.Currency{}
	currencyID := chi.URLParam(r, "currency_id")
	currencyIDColumn := fmt.Sprintf("%s.id", models.TableCoreCurrencies)
	condition := builder.Equal(currencyIDColumn, currencyID)

	return load(r, &currency, "LoadCurrency", models.TableCoreCurrencies, condition)
}

// UpdateCurrency updates object data in the database
func UpdateCurrency(r *http.Request) *Response {
	currencyID := chi.URLParam(r, "currency_id")
	currencyIDColumn := fmt.Sprintf("%s.id", models.TableCoreCurrencies)
	condition := builder.Equal(currencyIDColumn, currencyID)
	currency := models.Currency{
		ID: currencyID,
	}

	return update(r, &currency, "UpdateCurrency", models.TableCoreCurrencies, condition)
}

// DeleteCurrency deletes object from the database
func DeleteCurrency(r *http.Request) *Response {
	currencyID := chi.URLParam(r, "currency_id")
	currencyIDColumn := fmt.Sprintf("%s.id", models.TableCoreCurrencies)
	condition := builder.Equal(currencyIDColumn, currencyID)

	return remove(r, "DeleteCurrency", models.TableCoreCurrencies, condition)
}
