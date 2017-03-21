package main

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"net/http"
	"time"
)

func fetch() {
	// r, err := http.Get("https://www.reddit.com/r/dadjokes.json")
	// if err != nil {
	// 	// handle error
	// 	log.Fatal(err)
	// } else {
	// 	defer r.Body.Close()
	// 		_, err := io.Copy(os.Stdout, r.Body)
	// 		if err != nil {
	// 			// handle error
	// 			log.Fatal(err)
	// 		}
	// }

	client := &http.Client{}
	url := "http://facebook.com"

	// resp, err := client.Get(url)
	resp, err := http.NewRequest("GET", url, nil)
	startRes := time.Now()
	response, err := client.Do(resp)

	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		// _, err := io.Copy(os.Stdout, response.Body)
		if err != nil {
			log.Fatal(err)
		}

		endRes := time.Now()
		fmt.Println(" ")
		color.Yellow("Response %s", endRes.Sub(startRes))
	}

}

func main() {
	fetch()
}
