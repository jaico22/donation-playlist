package users

import (
	"internal/database"
	"log"

	"fmt"
	"net/http"
)

func NewUser(user NewUserRequest) NewUserResponce {
	fmt.Println("Registering New User")
	responce := NewUserResponce{}
	validateNewUserRequest(user, responce)
	writeUser(user)
	responce.StatusCode = 200
	return responce
}

func GetUser(userId string) GetUserResponce {
	var resp GetUserResponce
	user := database.GetUser(userId)
	if user == nil {
		resp.StatusCode = http.StatusNotFound
		return resp
	}
	resp.FacebookUserId = user.FacebookUserId
	resp.StatusCode = http.StatusOK
	return resp
}

func validateNewUserRequest(user NewUserRequest, responce NewUserResponce) NewUserResponce {
	if len(user.FacebookUserId) <= 0 {
		responce = invalidateResponce("FacebookUserId Required", responce)
		return responce
	}
	return responce
}

func invalidateResponce(message string, responce NewUserResponce) NewUserResponce {
	responce.ErrorMessage = message
	responce.StatusCode = http.StatusBadRequest
	return responce
}

func writeUser(user NewUserRequest) {
	log.Printf("UserId: %v", user.FacebookUserId)
	database.UpsertUser(database.UserDto{FacebookUserId: user.FacebookUserId})
}
