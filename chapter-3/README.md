# HTTP Clients and Remote Interaction with Tools - Chapter 3

- Will build a HTTP client and consume APIs of Shodan, Bing, and Metasploit 
- Will search and parse document metadata in a manner similar to the metadata search tool FOCA. 

## HTTP Fundamentals with Go

- HTTP is a stateless protocol; the server doesn't inherently maintain state and status for each request.
- State is tracked through a variety of means which may include session identifiers, cookies, HTTP headers, and more.

- Communication between server and client can be synchronously or asynchronously, but they operate on a request/response cycle.

- APIs commonly communicate via more structured data encoding, such as XML, JSON or MSGRPC. In some cases the data may be binary format, representing an arbitrary file type for download.

### Calling HTTP APIs

- Go's net/http standard package contains serveral convenience functions to quickly and easily send POST, GET, and HEAD request. 

```
Get(url string) (resp *Response, err error)
Head(url string) (resp *Response, err error)
Post(url string, bodyType string, body io.Reader) (resp *Response, err error)
```

- Post takes two additional parameters: bodyType, which is a string value that you use for the Content-Type HTTP header (commonly application/x-www-form-urlencoded) of the  request body, and an io.Reader

- Go has an additional POST request convenience functions, called PostForm(), which removes tediousness of setting those values and manually encoding every request; 

```
func PostForm(url string, data url.Values) (resp *Response, err error)
```

### Generating a Request

- To generate a request with one of the verbs (POST, GET, DELETE, PATCH, PUT), you can use the NewRequest() function to create the Request struct, which will be sent using the Client function's Do() method. (http.NewRequest())

```
func NewRequest(method, url string, body io.Reader) (req *Request, err error)
```

- A DELETE request

```
req, err := http.NewRequest("DELETE", "https://www.google.com/robots.txt", nil)
var client http.Client
resp, err := client.Do(req)
// Read response body and close.
```

### Using Structured Response Parsing

- Inspecting various components of the HTTP response is a crucial aspect of any HTTP-related task, like reading the response body, accessing cookies and headers, or simply inspecting the HTTP status code. 

- The Response type contains an exported Body parameter, which is of type io.ReadCloser.

- An io.ReadCloser is an interface that requires the contract of a Close() function and an io.Reader

## Building an HTTP Client That Interacts with Shodan

- Before performing any authorized adversarial activies against an organization, begin with reconnaissance.

- This starts with passive techniques that don't send packets to the target; that way, detection of the activity is next to impossible.

- Attackers use a variety of sources and services - including social networks, public records, and search engines to gain potentially useful information about the target.

- Shodan (https://www.shodan.io/), self-described as "the world's first search engine for internet-connected devices"

- Think of Shodan as a repository of scan data, even if it does much, much more.

### Reviewing the Steps for Building an API Client

- Building an API Client that interacts with Shodan API, parsing the results and displaying relevant information.

### Designing the Project Structure

- The project will be structured in a resuable way, allowing the packages to be as a library in its own packages.

```
package shodan
\\ shodan
\\\\api.go
\\\\host.go
\\\\shodan.go
```

### Cleaning Up Api Calls

- Building a resuable stack to rebuild the http request headers/host from.

```
func APIInfo(token, url string) { --snip-- }

func HostSearch(token, url string) { --snip-- }
```

### Querying Your Shodan Subscription

- Interation with Shodan, per the documentation, the call to query your subscription plan information as follow:

```
https://api.shodan.io/api-info?key={YOUR_API_KEY}
```

- Building the types that will allow us to Unmarshal the response.

### Creating a Client

- Implementing a client to take a search term from the command line and using the created API to interface with shodan services..

## Interacting with Metasploit

- Metasploit is a frame work used to perform a variety of adversarial techniques, including reconnaissance, exploitation, command and control, persistence, lateral network movement, payload creation and delivery, privilege escalation, and more. 

### Setting Up The Environment

- Download and install the Metasploit community edition. 

- Starting the Metasploit and the msgrpc server

```
msfconsole

load msgrpc Pass=s3cr3t ServerHost=10.0.1.6
```