package main

import (
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"./src/api"
)

func main() {
	router := httprouter.New()
	router.POST("/", api.PostImage)
	router.GET("/static/images/:name", api.GetImage)
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		buf, _ := ioutil.ReadFile("./index.html")
		w.Header().Set("content-type", "text/html")
		w.Write(buf)
	})

	log.Fatal(http.ListenAndServe(":10991", router))
}