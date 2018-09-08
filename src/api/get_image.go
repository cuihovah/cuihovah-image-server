package api

import (
	"github.com/julienschmidt/httprouter"
	"gopkg.in/h2non/filetype.v1"
	"io/ioutil"
	"log"
	"encoding/json"
	"net/http"
)

func GetImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	quality := r.URL.Query().Get("quality")
	if quality != "low" {
		quality = "images"
	}
	content, err := ioutil.ReadFile("./static/" + quality + "/" + ps.ByName("name"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		retval, _ := json.Marshal(Return{4, "service error", err.Error()})
		w.Write(retval)
	}
	kind, unknow := filetype.Match(content)
	if unknow != nil {
		log.Fatal(unknow.Error())
	}
	w.Header().Set("Content-Type", kind.MIME.Value)
	w.Write(content)
}