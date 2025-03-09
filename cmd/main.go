package main

import (
	"fmt"
	"log"
	"os"

	interurl "github.com/abhinaenae/crawli/pkg/url"
)

func main() {
	if len(os.Args[1:]) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	if len(os.Args[1:]) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	rawBaseURL := os.Args[1]
	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)

	htmlBody, err := interurl.GetHTML(rawBaseURL)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println(htmlBody)
}
