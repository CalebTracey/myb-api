package endpoints

import (
	"encoding/json"
	"github.com/calebtracey/mind-your-business-api/external"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// NewUser route handler for /newUser endpoint
//
// @Summary      New User request
// @Description  request to add new user to the database
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  http.HandlerFunc
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /api/v1/newUser [post]
func (r *Router) NewUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		var apiResponse *external.Response

		sw := time.Now()
		rw.Header().Set("Content-Type", "application/json")

		if apiResponse = r.Service.NewUser(req.Context(), req.Body); apiResponse.Message.ErrorLog != nil {
			log.Errorf("/newUser - %v", apiResponse.Message.ErrorLog)
		}

		apiResponse.Message.AddMessageDetails(sw)
		rw.WriteHeader(apiResponse.Message.ErrorLog.GetHTTPStatus(len(apiResponse.Details)))
		_ = json.NewEncoder(rw).Encode(apiResponse)
	}
}
