package endpoints

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func (r *Router) Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.WithTime(time.Now().Local()).Infoln("still alive?")

		w.WriteHeader(http.StatusOK)
	}
}
