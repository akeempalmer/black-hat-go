package main

import (
	"fmt"
	"log"
	"os"

	"github.com/akeempalmer/black-hat-go/chapter-3/shodan/shodan"
	"github.com/joho/godotenv"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: shodan searchterm")
	}

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	apiKey := os.Getenv("SHODAN_API_KEY")
	s := shodan.New(apiKey)

	info, err := s.APIInfo()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Printf(
		"Query Credits: %d\nScan Credits: %d\n\n",
		info.QueryCredits, info.ScanCredits,
	)

	hostSearch, err := s.HostSearch(os.Args[1])
	if err != nil {
		log.Panicln(err)
	}

	for _, host := range hostSearch.Matches {
		fmt.Printf("%18s%8d\n", host.IPString, host.Port)
	}
}
