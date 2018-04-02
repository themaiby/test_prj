package helper

import (
	"encoding/json"
	log "github.com/kataras/golog"
)

// generate JSON response
type responseTemplate struct {
	Status struct {
		Success      bool
		Message      interface{}
		ResponseTime float64
	} `json:"Status"`
	ResponseBody interface{}
}
type jsonResponse struct{}

var JSONResponse jsonResponse

func (this *jsonResponse) Success(model interface{}, duration float64) []byte {
	// generate status, error and response time
	prepareStruct := responseTemplate{
		Status: struct {
			Success      bool
			Message      interface{}
			ResponseTime float64
		}{
			Success:      true,
			Message:      nil,
			ResponseTime: duration},
		ResponseBody: model,
	}

	response, err := json.Marshal(prepareStruct)

	if err != nil {
		log.Warn(err)
	}
	return response
}

func (this *jsonResponse) Error(message string, duration float64) []byte {
	// generate status, error and response time
	prepareStruct := responseTemplate{
		Status: struct {
			Success      bool
			Message      interface{}
			ResponseTime float64
		}{Success: false,
			Message:      message,
			ResponseTime: duration},
	}

	response, err := json.Marshal(prepareStruct)

	if err != nil {
		log.Warn(err)
	}
	return response
}
