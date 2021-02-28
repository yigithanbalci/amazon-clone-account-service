package service

import (
	"log"
	"net/http"
)

// StartWebServer adds routes to http Handle
func StartWebServer(port string) {
	log.Println("Starting http service at " + port)

	r := NewRouter()
	http.Handle("/", r)
	err := http.ListenAndServe(":"+port, r)

	if err != nil {
		log.Println("An error occured starting http listener at port: " + port)
		log.Println("Error: " + err.Error())
	}
}
