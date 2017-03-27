package main

import (
	"flag"
	"github.com/fatih/color"
	"log"
	"net/http"
	"time"
)

func fetch() {
	wordPtr := flag.String("url", "http://facebook.com", "a string")
	flag.Parse()
	url := *wordPtr

	client := &http.Client{}
	color.Yellow("Pinging %s", url)

	resp, err := http.NewRequest("GET", url, nil)
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
	fetch()
}
