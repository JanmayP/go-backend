package controllers

import (
	"fmt"
	"net/http"
)

// ROUTE: /sample-get
func HandleSampleGet(w http.ResponseWriter, r *http.Request) {

}

// ROUTE: /sample-post
type SamplePostReq struct {
	Name string `json:"name"`
}

func (req *SamplePostReq) Validate() error {
	if req.Name == "" {
		return fmt.Errorf("missing parameter from request body: name")
	}

	return nil
}

func HandleSamplePost(w http.ResponseWriter, r *http.Request, req *SamplePostReq) {
	WriteJson(w, http.StatusOK, req)
	return
}
