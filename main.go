package main

import (
	//"fmt"
	"net/http"

	"github.com/bulirma/syllables/controllers"
	"github.com/bulirma/syllables/services"
	"github.com/bulirma/syllables/config"

	"github.com/gorilla/mux"
)

func main() {
	tokenMgr := services.TokenManager {
		DbFile: config.TokenFile,
	}
	regCtr := controllers.NewRegistrationController(tokenMgr)

	router := mux.NewRouter()

	router.PathPrefix("/css/").Handler(
		http.StripPrefix("/css/", http.FileServer(http.Dir("./static/css/"))))
	router.PathPrefix("/images/").Handler(
		http.StripPrefix("/images/", http.FileServer(http.Dir("./static/images/"))))

	router.HandleFunc("/", controllers.ServeStaticFile("index.html")).Methods("GET")
	router.HandleFunc("/registration-complete",
		controllers.ServeStaticFile("registration-complete.html")).Methods("GET")
	router.HandleFunc("/registration", regCtr.RegistrationFormGet).Methods("GET")
	router.HandleFunc("/registration", regCtr.RegistrationFormPost).Methods("POST")

	http.ListenAndServe(":8080", router)
}
