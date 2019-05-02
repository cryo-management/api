package resources

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/andreluzz/go-sql-builder/builder"
	"github.com/andreluzz/go-sql-builder/db"
	"github.com/cryo-management/api/models"
	"github.com/cryo-management/api/services"
	"github.com/dgrijalva/jwt-go"
)

// Login validate credentials and return user token
func Login(r *http.Request) *services.Response {
	response := services.NewResponse()
	body, _ := ioutil.ReadAll(r.Body)

	jsonMap := make(map[string]interface{})
	err := json.Unmarshal(body, &jsonMap)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, services.NewResponseError(services.ErrorParsingRequest, "Login unmarshal body", err.Error()))
		return response
	}

	_, emailOk := jsonMap["email"]
	_, passwordOk := jsonMap["password"]
	if !emailOk || !passwordOk {
		err = errors.New("Invalid credentials body")
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, services.NewResponseError(services.ErrorParsingRequest, "Login parsing body", err.Error()))
		return response
	}

	user := models.User{}
	emailColumn := fmt.Sprintf("%s.email", models.TableCoreUsers)
	err = db.LoadStruct(models.TableCoreUsers, &user, builder.Equal(emailColumn, jsonMap["email"]))
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, services.NewResponseError(services.ErrorLoadingData, "Login load user", err.Error()))

		return response
	}

	if user.ID == "" {
		err = errors.New("Invalid user email")
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, services.NewResponseError(services.ErrorLogin, "Login validation", err.Error()))
		return response
	}

	if user.Password != jsonMap["password"].(string) {
		err = errors.New("Invalid user password")
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, services.NewResponseError(services.ErrorLogin, "Login validation", err.Error()))
		return response
	}

	user.Password = ""
	claims := models.UserCustomClaims{
		user,
		jwt.StandardClaims{
			Issuer: "API",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	cryoSigningKey := []byte("AllYourBase") // TODO: Check the best place for this key, probably the config.toml
	tokenString, err := token.SignedString(cryoSigningKey)

	jsonMap["user"] = user
	jsonMap["token"] = tokenString
	delete(jsonMap, "password")
	delete(jsonMap, "email")
	response.Data = jsonMap

	return response
}
