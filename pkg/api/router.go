package api

import (
	"backend/pkg/api/controllers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (api *API) InitRouter() (*chi.Mux, error) {
	fmt.Println("initialising router")

	r := chi.NewRouter()

	// Initialise Middlewares
	r.Use(middleware.Logger)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/check", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("okay"))
		})

		r.Get("/sample-get", BaseMiddleware(PermissionMiddleware(controllers.HandleSampleGet)))
		r.Post("/sample-post", BaseMiddleware(PermissionMiddleware(ValidationMiddleware(controllers.HandleSamplePost))))
	})

	fmt.Println("registered routes")

	return r, nil
}

// can split routes and mount them
