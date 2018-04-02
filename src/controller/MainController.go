package controller

import (
	log "github.com/kataras/golog"
	"net/http"

	"../helper"
	"../model"
	"github.com/gorilla/mux"
	"strconv"
)

type mainController struct{}

var MainContoller = new(mainController)

// controller methods
func (this *mainController) MainPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Error("[mux VAR] Error: ", err)
	}

	result := model.User.Find(userID)

	//jsonRes, err := json.Marshal(result)
	if err != nil {
		log.Error("[JSON] Error: ", err)
	}

	duration := helper.ResponseTime.Get(r)

	// calculate request Time

	w.Header().Add("Content-type", "application/json")
	// we can't have id=0, so return error 404
	if result.ID == 0 {
		w.Write(JSONResponse.Error("User does not exists", duration))
		return
	}
	w.Write(JSONResponse.Success(result, duration))
}
