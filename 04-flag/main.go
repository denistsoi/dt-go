package main

import (
	"flag"
	"github.com/fatih/color"
	"log"
	"net/http"
	"net/url"
	"time"
)

func fetch(URL string) {
	client := &http.Client{}
	color.Yellow("Pinging %s", URL)

	resp, err := http.NewRequest("GET", URL, nil)
	startRes := time.Now()
	response, err := client.Do(resp)

	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		if err != nil {
			log.Fatal(err)
		}

		endRes := time.Now()
		switch response.StatusCode {
		case http.StatusOK:
			color.Green("Response Ok! 👌")
		default:
			color.Red("Error %d", response.StatusCode)
		}
		color.Yellow("Response %s", endRes.Sub(startRes))
	}
}

func main() {
	wordPtr := flag.String("url", "http://facebook.com", "a string")
	flag.Parse()

	URL, err := url.Parse(*wordPtr)
	if err != nil {
		log.Fatal(err)
	}

	URLSTRING := URL.String()

	if URL.Scheme == "" {
		log.Print("Scheme is missing")
		URLSTRING = "http://" + URL.String()
	}
	fetch(URLSTRING)
}
