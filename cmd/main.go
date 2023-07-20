package main

import (
	"backend/pkg/api"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/oklog/run"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGSEGV)

	// make core run group
	g := run.Group{}

	api, err := api.InitAPIServer()
	if err != nil {
		fmt.Printf("could not init API server: %v\n", err)
	}

	g.Add(func() error {
		return api.ListenAndServe()
	}, func(error) {
		_ = api.Close()
	})

	go func() {
		fmt.Println("running core group")

		err = g.Run()
		if err != nil {
			// shutdown logic
		}
	}()

	gracefulShutdown(sigs, api)
}

func gracefulShutdown(sigs chan os.Signal, api *api.API) {
	sig := <-sigs
	fmt.Printf("received shutdown signal: %v\n", sig)

	// cleanup logic here
	fmt.Println("cleaning up...")

	_ = api.Close()
}
