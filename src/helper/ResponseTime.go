package helper

import (
	"github.com/gorilla/context"
	"net/http"
	"time"
)

type responseTime struct{}

func (this *responseTime) Get(r *http.Request) float64 {

	val, _ := context.GetOk(r, "RequestTime")
	responseTime := time.Now().Sub(val.(time.Time))

	return responseTime.Seconds()
}

var ResponseTime responseTime
