package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		fmt.Println(address)

		conn, err := net.Dial("tcp", address)

		if err != nil {
			fmt.Printf("Port connection failed: %s\n", address)
			continue
		}

		conn.Close()
		fmt.Printf("%d open\n", i)
	}
}
