package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// repo <- service -> serializer  -> http

var port string
var mux *http.ServeMux
var fsWeb http.Handler

func main() {
	loadEnv()

	mux.Handle("/web/html/", http.StripPrefix("/web/html/", fsWeb))
	log.Println("Listing for requests at http://localhost:8000/index")
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		// log.Fatal("[ERROR]: " + err.Error())
	}

	port = os.Getenv("PORT")
	if port == "" {
		port = "8000"
		//log.Fatal("[ERROR]: Port has not been set in .env file!")
	}

	mux = http.NewServeMux()
	fsWeb = http.FileServer(http.Dir("web/html"))
}
