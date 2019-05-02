package resources

import (
	"fmt"
	"net/http"

	"github.com/andreluzz/go-sql-builder/builder"
	"github.com/go-chi/chi"

	"github.com/cryo-management/api/models"
	"github.com/cryo-management/api/services"
)

// CreateCurrency persists the request body creating a new object in the database
func CreateCurrency(r *http.Request) *services.Response {
	currency := models.Currency{}

	return services.Create(r, &currency, "CreateCurrency", models.TableCoreCurrencies)
}

// LoadAllCurrencies return all instances from the object
func LoadAllCurrencies(r *http.Request) *services.Response {
	currencies := []models.Currency{}

	return services.Load(r, &currencies, "LoadAllCurrencies", models.TableCoreCurrencies, nil)
}

// LoadCurrency return only one object from the database
func LoadCurrency(r *http.Request) *services.Response {
	currency := models.Currency{}
	currencyID := chi.URLParam(r, "currency_id")
	currencyIDColumn := fmt.Sprintf("%s.id", models.TableCoreCurrencies)
	condition := builder.Equal(currencyIDColumn, currencyID)

	return services.Load(r, &currency, "LoadCurrency", models.TableCoreCurrencies, condition)
}

// UpdateCurrency updates object data in the database
func UpdateCurrency(r *http.Request) *services.Response {
	currencyID := chi.URLParam(r, "currency_id")
	currencyIDColumn := fmt.Sprintf("%s.id", models.TableCoreCurrencies)
	condition := builder.Equal(currencyIDColumn, currencyID)
	currency := models.Currency{
		ID: currencyID,
	}

	return services.Update(r, &currency, "UpdateCurrency", models.TableCoreCurrencies, condition)
}

// DeleteCurrency deletes object from the database
func DeleteCurrency(r *http.Request) *services.Response {
	currencyID := chi.URLParam(r, "currency_id")
	currencyIDColumn := fmt.Sprintf("%s.id", models.TableCoreCurrencies)
	condition := builder.Equal(currencyIDColumn, currencyID)

	return services.Remove(r, "DeleteCurrency", models.TableCoreCurrencies, condition)
}

// CreateCurrencyRate persists the request body creating a new object in the database
func CreateCurrencyRate(r *http.Request) *services.Response {
	currencyRate := models.CurrencyRate{}

	return services.Create(r, &currencyRate, "CreateCurrencyRate", models.TableCoreCryRates)
}

// LoadAllCurrencyRates return all instances from the object
func LoadAllCurrencyRates(r *http.Request) *services.Response {
	currencies := []models.CurrencyRate{}
	fromCurrencyCode := chi.URLParam(r, "currency_code")
	fromCurrencyCodeColumn := fmt.Sprintf("%s.from_currency_code", models.TableCoreCryRates)
	condition := builder.Equal(fromCurrencyCodeColumn, fromCurrencyCode)

	return services.Load(r, &currencies, "LoadAllCurrencyRates", models.TableCoreCryRates, condition)
}

// LoadCurrencyRate return only one object from the database
func LoadCurrencyRate(r *http.Request) *services.Response {
	currencyRate := models.CurrencyRate{}
	currencyRateID := chi.URLParam(r, "currency_rate_id")
	currencyRateIDColumn := fmt.Sprintf("%s.id", models.TableCoreCryRates)
	condition := builder.Equal(currencyRateIDColumn, currencyRateID)

	return services.Load(r, &currencyRate, "LoadCurrencyRate", models.TableCoreCryRates, condition)
}

// UpdateCurrencyRate updates object data in the database
func UpdateCurrencyRate(r *http.Request) *services.Response {
	currencyRateID := chi.URLParam(r, "currency_rate_id")
	currencyRateIDColumn := fmt.Sprintf("%s.id", models.TableCoreCryRates)
	condition := builder.Equal(currencyRateIDColumn, currencyRateID)
	currencyRate := models.CurrencyRate{
		ID: currencyRateID,
	}

	return services.Update(r, &currencyRate, "UpdateCurrencyRate", models.TableCoreCryRates, condition)
}

// DeleteCurrencyRate deletes object from the database
func DeleteCurrencyRate(r *http.Request) *services.Response {
	currencyRateID := chi.URLParam(r, "currency_rate_id")
	currencyRateIDColumn := fmt.Sprintf("%s.id", models.TableCoreCryRates)
	condition := builder.Equal(currencyRateIDColumn, currencyRateID)

	return services.Remove(r, "DeleteCurrencyRate", models.TableCoreCryRates, condition)
}
