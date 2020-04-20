package main

import (
	"fmt"
	"net/http"
	"os"
	router "template/Routers"

	"github.com/gorilla/handlers"
)

func main() {
	r := router.InitializeRouter()
	fmt.Println(" your server is runing in port 8000 ")
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	err := http.ListenAndServe(":"+port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"https://signfront.herokuapp.com", "http://localhost:3000"}))(r))
	fmt.Println(err)
}
