package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	helper "github.com/shwetank0714/jwtmodfile/helpers"
	model "github.com/shwetank0714/jwtmodfile/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var createdUserInput model.CreateUser

	json.NewDecoder(r.Body).Decode(&createdUserInput)

	getCreatedUser := helper.CreateUser(&createdUserInput)

	json.NewEncoder(w).Encode(getCreatedUser)

	defer r.Body.Close()
}

func GetAllUsersList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allUsersList := helper.GetAllUsers()

	json.NewEncoder(w).Encode(allUsersList)
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var loginCred model.LoginModel
	// w.WriteHeader(http.StatusBadGateway)

	json.NewDecoder(r.Body).Decode(&loginCred)

	emailId := loginCred.Email
	password := loginCred.Password

	userModel, err := helper.UserLoginHelper(emailId, password)

	if err != nil {

		json.NewEncoder(w).Encode(err.Error())
	} else {
		json.NewEncoder(w).Encode(userModel)
	}

}

func ChangePassword(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	token := r.Header.Get("Authorization")
	log.Println(token)
	if token == "" {
		log.Fatal(http.StatusBadRequest)
	}
	var changePasswordRequest model.ChangePasswordModel

	json.NewDecoder(r.Body).Decode(&changePasswordRequest)

	updatedUserData, err := helper.ChangeUserPassword(changePasswordRequest, token)
	if err != nil {
		log.Fatal(http.StatusUnauthorized)
	}
	json.NewEncoder(w).Encode(updatedUserData)
}
