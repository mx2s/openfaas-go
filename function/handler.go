package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := getRouter()

	log.Fatal(http.ListenAndServe(":8201", router))
}

func getRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", index)
	router.POST("/post-route", postHandler)

	return router
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Hello from julienschmidt/httprouter !")
}

func postHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "POST works too!")
}
