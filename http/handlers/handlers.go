package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

// HealthCheck responds with a server ok during server health check
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}
func PostHealthCheck(w http.ResponseWriter, r *http.Request) {
	params := Parameters{}
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		//	Log to the system and return
		apiResponse := Response{
			Status:  "failed",
			Comment: "Failed to read JSON Body",
		}
		// Send response
		render.JSON(w, r, apiResponse)

		log.Error("Failed to read JSON Body")
		return
	}
	err = json.Unmarshal(body, &params)
	//	Log to the system and return
	response := Response{
		Status:  "ok",
		Comment: "All is good here",
	}
	// Send response
	render.JSON(w, r, response)
}