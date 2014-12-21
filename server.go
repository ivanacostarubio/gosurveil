/*

Sever Library Gather Data from Clients

Install


TODO: Create machine filder only if not present

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
	"path/filepath"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Log struct {
	Hostname string `param:"hostname"`
	Message  string `param:"message"`
}

func createDirectoryForHost(host string) {

	os.Mkdir("."+string(filepath.Separator)+"/tmp/"+host, 0777)
}

func timeBasedName(host string) string {

	createDirectoryForHost(host)

	t := time.Now()
	temp := "./tmp/"
	fileName := "20060102150405.png"
	r := temp + host + "/" + fileName

	return t.Format(r)

}

func writeToFile(data string, host string) {

	f, err := os.Create(timeBasedName(host))
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

	fmt.Printf("El Greeat.Message: %s\n", log.Hostname)
	d1 := []byte(log.Message)

	host := strings.TrimSpace(log.Hostname)

	if len(d1) > 1 {
		writeToFile(string(d1), host)
		w.WriteHeader(201)
		return
	}

	if debug {
		fmt.Printf("El Greeat.Message: %s\n", log.Hostname)
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
