package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/vahdet/go-auth-service/models"
	"github.com/vahdet/go-auth-service/services/interfaces"
)

type (
	AuthController struct {
		service interfaces.AuthService
	}
)

func NewAuthController(service interfaces.AuthService) *AuthController {
	return &AuthController{service}
}

func (ac *AuthController) Register(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Stub an user to be populated from the body
	u := models.User{}

	// Populate the user data from the request body
	json.NewDecoder(r.Body).Decode(&u)

	// do service call
	createdUser, _ := ac.service.CreateUser(&u)
	//TODO: Handle error

	uj, _ := json.Marshal(createdUser)
	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)

}

func (ac *AuthController) Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Stub an user to be populated from the body
	u := models.User{}

	// Populate the user data from the request body
	json.NewDecoder(r.Body).Decode(&u)

	// Get user from path variable
	userParam := p.ByName("user")

	// Check if the request is made through username or email
	e, err := mail.ParseAddress(userParam)
	if err != nil {
		u.Name = userParam
	} else {
		u.Email = e.Address
	}

	userId := int64(1234)

	//TODO: IMPLEMENT PASSWORD CHECK LOGIC HERE OR IN SERVICE
	// do service call

	createdUser, _ := ac.service.GetTokens(userId)
	//TODO: Handle error

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(createdUser)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

func (ac *AuthController) Logout(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Get user id from path variable
	userId := p.ByName("userid")

	userIdInt64, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		w.WriteHeader(500)
	}

	ac.service.PurgeRefreshToken(userIdInt64)
	//TODO: Handle error

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	// fmt.Fprintf(w, "%s", userId)
}
