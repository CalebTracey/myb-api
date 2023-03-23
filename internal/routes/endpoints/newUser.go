package endpoints

import (
	"encoding/json"
	"github.com/calebtracey/mind-your-business-api/external"
	"github.com/calebtracey/models/pkg/response"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func (r *Router) NewUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		sw := time.Now()
		apiRequest := new(external.ApiRequest)

		if err := json.NewDecoder(req.Body).Decode(apiRequest); err != nil {
			log.Errorf("NewUser: error: %v", err)
		}

		apiResponse := r.Service.NewUser(req.Context(), apiRequest)
		statusCode := apiResponse.Message.ErrorLog.GetHTTPStatus(len(apiResponse.Details))
		apiResponse.Message.AddMessageDetails(sw)

		if res, err := json.Marshal(apiResponse); err != nil {
			log.Errorf("failed to marshal response; error: %s", err.Error())
			statusCode = http.StatusInternalServerError
		} else {
			response.WriteHeader(rw, statusCode)
			_, _ = rw.Write(res)
		}
	}
}
