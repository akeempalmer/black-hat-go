package main

import (
	"fmt"
	"log"
	"os"

	"github.com/akeempalmer/black-hat-go/rpc"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	host := os.Getenv("MSFHOST")
	pass := os.Getenv("MSFPASS")
	user := os.Getenv("MSFUSER")

	if host == "" || pass == "" {
		log.Fatalln("Missing required environment variable MSFHOST or MSFPASS")
	}
	msf, err := rpc.New(host, user, pass)
	if err != nil {
		log.Panicln(err)
	}

	defer msf.Logout()

	sessions, err := msf.SessionList()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("sessions:")
	for _, session := range sessions {
		fmt.Printf("%5d %s\n", session.ID, session.Info)
	}
}
