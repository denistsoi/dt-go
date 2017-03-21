package main

import (
	"log"
	"net/http"
	"time"

	"github.com/fatih/color"
)

func fetch() {
	client := &http.Client{}
	url := "http://facebook.com"
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
