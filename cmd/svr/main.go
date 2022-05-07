package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"serverTemplate/internal/handler"
)

func main() {
	router := handler.InitializeRoutes(new(mux.Router))
	fmt.Println("App Listening On PORT:8000")
	log.Fatal(http.ListenAndServe(":8000", &router))
}
