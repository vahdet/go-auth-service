package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

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

func (ac AuthController) Register(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Stub an user to be populated from the body
	u := models.User{}

	// Populate the user data
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

func (ac AuthController) Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Stub an example user
	u := models.User{
		Name:   "Bob Smith",
		Gender: "male",
		Age:    50,
		Id:     p.ByName("id"),
	}
	//TODO: IMPLEMENT PASSWORD CHECK LOGIC HERE OR IN SERVICE
	// do service call
	createdUser, _ := ac.service.GetTokens(userId)
	//TODO: Handle error

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

func (ac AuthController) Logout(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ac.service.CreateUser(&u)
}
