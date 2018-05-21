package main

import (
	"fmt"
	"log"
	"net/http"

	ghandler "github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"microservice_boilerplate_httpRestLike/config"
)

func main() {

	// base server
	var router = mux.NewRouter()
	router.HandleFunc("/checkhealth", apiCheckHealth)

	// api routes
	//var routerApi = router.PathPrefix("/api").Subrouter()

	url := config.ADDRESS + config.PORT
	/*
		we use CORS to allow call from xml http request of client's app by example
	*/
	headersOk := ghandler.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := ghandler.AllowedOrigins([]string{"*"})
	methodsOk := ghandler.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	/*
		Run server by using specified config
	*/
	if config.UseTls {
		fmt.Println("launching app on url: " + "https://" + "127.0.0.1" + config.PORT)
		var err error
		if config.UseCors {
			err = http.ListenAndServeTLS(url, config.TlsFileCert, config.TlsFileKey, ghandler.CORS(originsOk, headersOk, methodsOk)(router))
		} else {
			err = http.ListenAndServeTLS(url, config.TlsFileCert, config.TlsFileKey, router)
		}
		if err != nil {
			log.Fatal("listen and serve: ", err)
		}
	} else {
		fmt.Println("launching app on url: " + "http://" + "127.0.0.1" + config.PORT)
		var err error
		if config.UseCors {
			err = http.ListenAndServe(url, ghandler.CORS(originsOk, headersOk, methodsOk)(router))
		} else {
			err = http.ListenAndServe(url, router)
		}
		if err != nil {
			log.Fatal("listen and serve: ", err)
		}
	}
}

// simple http response
func apiCheckHealth(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("ok"))
}
