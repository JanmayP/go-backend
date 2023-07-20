package api

import (
	"backend/pkg/db"
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/rs/zerolog"
)

// can add a config initialiser for api (for dependency injection)
// type Config struct {
//  port string
//  aws aws
//  ihomer ihomer
// 	db  *sql.DB
// 	log zerolog.Logger
// }

type API struct {
	server *http.Server
	log    zerolog.Logger
	db     *sql.DB
}

func InitAPIServer() (*API, error) {
	fmt.Println("initialising api")

	api := &API{}

	// init router & server
	router, err := api.InitRouter()
	if err != nil {
		return nil, fmt.Errorf("could not init router: %v", err)
	}

	api.server = &http.Server{
		Addr:    ":8001",
		Handler: router,
	}

	// init db
	db, err := db.Init()
	if err != nil {
		return nil, fmt.Errorf("could not init db: %v", err)
	}

	api.db = db

	// init logger zzz
	api.log = zerolog.Logger{}

	return api, nil
}

func (api *API) ListenAndServe() error {
	fmt.Println("serving on port 8001")
	return api.server.ListenAndServe()
}

func (api *API) Close() error {
	fmt.Println("closing server")
	return api.server.Shutdown(context.Background())
}
