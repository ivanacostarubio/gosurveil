/*

Sever Library Gather Data from Clients

Install


Ubuntu:

apt-get install git
apt-get install golang
go get github.com/zenazn/goji
go get github.com/goji/param

*/

package main

import (
	"fmt"
	"github.com/goji/param"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Log struct {
	Numero  string `param:"numero"`
	Message string `param:"message"`
}

func timeBasedName() string {
	t := time.Now()
	return t.Format("./tmp/20060102150405.png")
}

func writeToFile(data string) {

	f, err := os.Create(timeBasedName())
	check(err)
	defer f.Close()

	d, err := f.WriteString(data)
	check(err)
	fmt.Println(d)
	f.Sync()

}

func log(c web.C, w http.ResponseWriter, r *http.Request) {

	var debug = true
	var log Log

	r.ParseForm()

	err := param.Parse(r.Form, &log)

	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	d1 := []byte(log.Message)

	if len(d1) > 1 {
		writeToFile(string(d1))
		w.WriteHeader(201)
		return
	}

	if debug {
		//  fmt.Printf("El Greeat.Message: %s\n", log.Message)
		//	fmt.Printf("El Great Completo: %s\n", log)
	}

	http.Error(w, http.StatusText(304), 304)
	return
}

func main() {
	goji.Post("/log/", log)
	staticFilesLocation := "./"
	goji.Handle("/*", http.FileServer(http.Dir(staticFilesLocation)))
	goji.Serve()
}
