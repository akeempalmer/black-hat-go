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