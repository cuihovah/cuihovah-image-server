package api

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"../image"
	"io/ioutil"
	"net/http"
)

func PostImage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
	image.MakeThumbnail("./static/images/"+fileName, "./static/low/"+fileName)
	fname = "/static/images/" + fileName
	retval, _ := json.Marshal(Return{0, "OK", fname})
	w.Header().Set("content-type", "application/json")
	w.Write(retval)
}