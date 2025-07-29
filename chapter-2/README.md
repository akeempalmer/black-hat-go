# TCP, Scanners, and Proxies - Chapter 2

- In this chapter we will learn the basic of TCP (Transmission Control Protocol)
- Build a concurrent, properly throttled port scanner
- Create a TCP proxy that can be used for port fowarding
- Re-create Netcat's gaping security hole feature

## Understanding the TCP Handshake

- If a port is open, a three-way handshake takes place. client -> syn -> Server; client <- syn-act  <- server; client -> ack -> server;

- If the port is closed, the server responds with a rst packet instead of a  syn-ack. Client -> syn -> Server; Client <- rst <- Server

- If the traffic is filtered by a firewall, the client will typically receive no response from the server.


## Bypassing Firewalls with Port Forwarding

- Using an intermediary system to proxy the connection around or through a firewall - port forwarding.

Client -> request stacktitan.com -> request traverse firewall -> stacktitan.com -> traffic proxied to evil.com -> evil.com


## Writing a TCP Scanner

- A port scanner may scan serveral ports by using a single contiguous method;, this can be time-consuming when your goal is to scall all 65, 535 ports. 

### Testing for Port Availability

- Using Go's net package method Dial(network string, address string) to communicate with Layer 4 networks 

- network - specifies the communication protocol to use "TCP|UDP|etc"

- address - specifies the host:port name "scanme.nmap.org:80"

- Dial returns conn and error results.

```
See ./dial/main.go

for code example
```

### Performing Nonconcurrent Scanning

- TCP ranges from 1 to 65535; for testing 1 to 1024.

- We conduct a loop from 1 to 1024, we need to convert the int to a string:

```
strconv

fmt.Sprintf(format string, a ...interfcae{})
```

### Performing Concurrent Scanning

- Scanning multiple ports concurrently. 

- Go allows you to create as many goroutines as your system can handle, bounded by available memory.

- seeing the time it took for the job to complete.

```
time ./tcp-scanner-too-fast 
```

- Using the wait group to control the concurrency

```
var wg sync.WaitGroup
```

### Synchronized Scanning Using WaitGroup

- This version uses a waitgroup connection to managae concurrent workers. 

- This version is still to fast, that the connection is not established successfully, on large counts of calls.


### Port Scanning Using a Worker Pool

-  Using a pool of goroutines to manage concurrent work being performed. 

```
// Create our worker function for processing work
func worker(ports chan int, wg *sync.WaitGroup) {
	for p := range ports {
		fmt.Println(p)
		wg.Done()
	}

}
```

### Multichannel Communication

- https://github.com/blackhat0go/bhg/blob/master/ch-2/scanner-port-format/

## Building a TCP Proxy

- Building a TCP server, as before we used the net package as a client.

- Will be building a proxy server for transfering data.

### Using io.Reader and io.Writer

- Cornerstone package for all input/output tasks 

```
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}
```

- Using the Reader interface, Writer interface

```
type FooReader struct {}
func (fooReader *FooReader) Read(p []byte) (int,error) {
	// Read some data from somewhere, anywhere.
	return len(dataReadFromSomewhere), nil
}
```

- Using the copy function to copy data from Reader to a Writer

```
func Copy(dst io.Writer, src io.Reader) (written int64, error)
```

### Creating the Echo Server

- Using the .net.Conn package to start a server and learning how to read and write data to and from a socket.

- .net.Conn is Go's steam-oriented network connection.

- Conn implements both Read and Write function, so this is an instance of both the Reader and Writer interfaces.

- We will use net.Listen to open a TCP listener on a specific port, once a client connect, the Accept method creates and returns a Conn object, used for sending and recieving data.

### Improving the Code by Creating a Buffered Listener

- Using the bufio package wrapping a Reader and Writer to create a buffered I/O mechanism.

- bufio.NewReader creates a buffered reader, which calls ReadString('\n') with a delimiter for where to stop reading '\n'

- bufio.NewWriter creates a buffered writer, which calls WriteString() with passes the data to the socket in conn.

- After writing the data, you must call writer.Flush() to have the data written to the writer.

### Proxying a TCP Client

- Using io.Copy() to send bytes from src to dst, in both and using dial and accept from the net package for communicating with the client we are forwarding too. 

### Replicating Netcat for Command Execution

- Netcat is the TCP/IP Swiss Army knife - a more flexible version of Telnet.

- This command creates a listening server on port 13337

```
nc -lp 13337 -e /bin/bash
```

- Using Go's os/exec package for running operating systems commands.

- Command(name string, arg ...string)

- Creating a custom writer to use for the cmd.Stdout, which will handle flushing the output as they are sent.

- This flusher takes in the Writer which will be passed on from conn

- We will also introduce the pipe function io.Pipe(), Go's synchronous, in-memory pipe that can be used for connecting Redaers and Writers:

```
func Pipe() (*PipeReader, *PipeWriter)
```

- Using the PipeReader and PipeWriter allows you to avoid having to explicitly flush the writer and synchronously connect stdout and the TCP connection. 

- This netcat-exec enables us a a server listener expecting a connection, we can use similar setup to enable this as a client. 

1. Establish a connection to a remote listneer via net.Dial(network, address string)
2. Initialize a Cmd via exec.Command(name string, arg ...string).
3. Redirect Stdin and Stdout properties to utilize the net.Conn object
4. Run the command.

### Summary

- Exploring Go as it relates to networking, I/O, and conncurrency. 

- Code based ["https://github.com/blackhat-go/bhg/blob/master/ch-2/netcat-exec/main.go"](Code Base)