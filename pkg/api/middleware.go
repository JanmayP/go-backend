package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"
)

type RequestBody interface {
	Validate() error
}

func ValidateRequest(req RequestBody) error {
	err := req.Validate()
	if err != nil {
		return err
	}

	return nil
}

func ValidationMiddleware[T RequestBody](next func(w http.ResponseWriter, r *http.Request, req T)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// parse request body
		req := new(T)
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "error unmarshaling request body", http.StatusInternalServerError)
			return
		}
		// validate request body
		err = ValidateRequest(*req)
		if err != nil {
			http.Error(w, fmt.Sprintf("invalid request: %v", err), http.StatusBadRequest)
			return
		}

		next(w, r, *req)
	}
}

func PermissionMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		next(w, r)
	}
}

func BaseMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer panicRecoveryHelper(r.URL.RawPath)

		next(w, r)
	}
}

func panicRecoveryHelper(route string) {
	if rec := recover(); rec != nil {
		panicErr := fmt.Sprintf("%v", rec)
		stackTrace := string(debug.Stack())

		// Log
		fmt.Printf("panic handling route: %v", route)
		fmt.Printf("error: %v", panicErr)
		fmt.Printf("Stack trace: \n%v\n", stackTrace)
	}
}
