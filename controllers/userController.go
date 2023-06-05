package controllers

import (
	"encoding/json"
	"net/http"

	model "github.com/shwetank0714/jwtmodfile/models"
	helper "github.com/shwetank0714/jwtmodfile/helpers"
)


func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","POST")

	var createdUserInput model.CreateUser

	json.NewDecoder(r.Body).Decode(&createdUserInput)


	getCreatedUser := helper.CreateUser(&createdUserInput)

	json.NewEncoder(w).Encode(getCreatedUser)

	defer r.Body.Close()
}

func GetAllUsersList(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allUsersList := helper.GetAllUsers()

	json.NewEncoder(w).Encode(allUsersList)
}
