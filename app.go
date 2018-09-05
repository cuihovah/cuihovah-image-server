package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/h2non/filetype.v1"
	"io/ioutil"
	"log"
	"net/http"
)

type Return struct {
	Code int `json:"code"`
	Msg string `json:"string"`
	Data interface{} `json:"data"`
}

func main() {
	router := httprouter.New()
	router.POST("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		file, _, err := r.FormFile("fileUpload")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fname := ""
		defer file.Close()
		rd := bufio.NewReader(file)
		p := make([]byte, 10 * 1024 * 1024)
		size, _ := rd.Read(p)
		hash := fmt.Sprintf("%x", md5.Sum(p))
		fileName := string(hash[:])
		ioutil.WriteFile("./static/images/"+fileName, p[:size], 0644)
		fname = "/static/images/" + fileName
		retval, _ := json.Marshal(Return{0, "OK", fname})
		w.Header().Set("content-type", "application/json")
		w.Write(retval)
	})

	router.GET("/static/images/:name", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		content, _ := ioutil.ReadFile("./static/images/" + ps.ByName("name"))
		kind, unknow := filetype.Match(content)
		if unknow != nil {
			log.Fatal(unknow.Error())
		}
		w.Header().Set("Content-Type", kind.MIME.Value)
		w.Write(content)
	})

	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		buf, _ := ioutil.ReadFile("./index.html")
		w.Header().Set("content-type", "text/html")
		w.Write(buf)
	})

	log.Fatal(http.ListenAndServe(":10991", router))
}