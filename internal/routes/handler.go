package routes

import (
	"github.com/calebtracey/mind-your-business-api/internal/routes/endpoints"
	"github.com/go-chi/chi/v5"
	"github.com/swaggo/http-swagger"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
)

//go:generate swagger generate spec -o ./swagger.json --scan-models --exclude-main --include=./
type Handler struct {
	Router endpoints.RouterI
}

func (h Handler) RouteHandler() *chi.Mux {
	r := chi.NewRouter()
	setMiddleware(r)

	// Health route handler for /health endpoint
	//
	// @Summary      Health check endpoint
	// @Description  request to check for 200 response
	// @Tags         util
	// @Success      200
	// @Router       /health [get]
	r.Get(endpoints.Health, h.Router.Health())

	r.Route(v1BasePath, func(r chi.Router) {
		// NewUser route handler for /newUser endpoint
		//
		// @Summary      New User request
		// @Description  request to add new user to the database
		// @Tags         users
		// @Accept       json
		// @Produce      json
		// @Success      200  {object}  external.Response
		// @Failure      400  {object}  external.Response
		// @Failure      404  {object}  external.Response
		// @Failure      500  {object}  external.Response
		// @Router       /api/v1/newUser [post]
		r.Post(endpoints.NewUser, h.Router.NewUser())
	})

	// serve swagger static page: http://localhost:6080/swagger/index.html
	r.Route(swaggerBasePath, func(r chi.Router) {
		r.Get(wildCard, httpSwagger.Handler(
			httpSwagger.URL(swaggerUiPath+swaggerDoc)), //The url pointing to API definition
		)
	})

	return r
}

const (
	wildCard        = "/*"
	v1BasePath      = "/api/v1"
	swaggerBasePath = "/swagger"
	swaggerDoc      = "doc.json"
	swaggerUiPath   = "http://localhost:6080/swagger/"
)
