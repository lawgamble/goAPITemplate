package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// InitializeRoutes - put all of your routes here
func InitializeRoutes(r *mux.Router) mux.Router {
	r.HandleFunc("/", handler)
	return *r
}

// handler - example of a route handler function
func handler(w http.ResponseWriter, r *http.Request) {
	var response map[string]interface{}
	_ = json.Unmarshal([]byte(`{ "hello": "world" }`), &response)
	respondWithJSON(w, http.StatusOK, response)
}

// respondWithJSON - called from handler func(s) - shows response/status
func respondWithJSON(w http.ResponseWriter, status int, res interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(res)
}
