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