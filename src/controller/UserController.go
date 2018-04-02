package controller

import (
	// go
	"net/http"
	"strconv"
	// packages
	"github.com/gorilla/mux"
	log "github.com/kataras/golog"
	// local
	"../helper"
	"../model"
)

type userController struct{}

var UserController = new(userController)

// controller methods
func (this *userController) GetUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Error("[mux VAR] Error: ", err)
	}

	result := model.User.Find(userID)

	if err != nil {
		log.Error("[JSON] Error: ", err)
	}

	duration := helper.ResponseTime.Get(r) // calculate request Time

	w.Header().Add("Content-type", "application/json")
	// we can't have id=0, so return error 404
	if result.ID == 0 {
		w.Write(helper.JSONResponse.Error("User does not exists", duration)) // TODO: make sql queries with err
		return
	}
	w.Write(helper.JSONResponse.Success(result, duration))
}

func (this *userController) GetAllUsers(w http.ResponseWriter, r *http.Request) {

}
