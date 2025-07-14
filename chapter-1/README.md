# Chapter 1

- This chapter covers the basic fundementals of go programming syntax, from for loops, to concurrency. These basic syntax would get us on track for the remaining of the book.

## Downloading and Installing Go

- We would download the Go binary realease from [https://golang.org/dl](Golang Download), selecting the binary for your specific computer OS.

## Setting the GOROOT to define the binaries location

- Set the reserved GOROOT enviornment variable to the location of the binary. On linux or mac, you can add this to your ~/profile.

```set GOROOT=/path/to/go```

- On windows, you can add this environment variable through the System (Control Panel), by clicking the Enviornment Variables button.

## Setting the GOPATH to determine the location of your go workspace

- Instructs the Go tool kit where your source code, third-party libs, and compiled programes will exist. For setting your projects in a directory call gocode in the home directory. Go project will contain three main directories, bin, pkg, annd src.

```GOPATH=$HOME/gocode```

- The bin directory hold all compiled executable, the pkg directory stores various package objects including third-party Go dependecies. The src directory will contain all the source code we will develop.

``` go version```

- Run the above for confirming that installation was successful.

## Choosing an IDE 

- [https://github.com/golang/go/wiki/IDEsAndTextEditorPlugins/]

## Vim Editor

- To use Vim for Go, install the vim-go plugin [https://github.com/faith/vim-go]

## Common Go Tools Commads

1. go run - will compile and execute the main package (programs entry point).
2. go build - produces a standalone binary file on disk, without execution. (can use -o to specifiy the output build's name).
    - By default the produced binary file contains debuggig information adn the symbol table, which can bloat the size of the file. To reduce the file size you ca include additional flags during the build process to strip this information from the binary.

```go build -ldflags "-w -s"```

3. Cross Compiling -  compiling for different architecture from your evniornment. [https://golang.org/doc/install/source#environment/]

- We need to pass the Constrants, which include GOOS (for the operating system) and GOARCH (for the architecture), This can be added in three ways 
1. Via Command line
2. Code comments
3. File suffix naming convention

- Building for Linux 64 bit 

```GOOS="linux" GOARCH="amd64" go build hello.go```

## The go doc Command

- Lets you interrogate documentation about a package, function, method, or variable. 

```go doc fmt.Println```

## The go get Command

- Lets you import package source code for third-party dependencies and places it within the $GOPATH/src/directory. 
- Uses the dep and mod tools to lock dependiencies in order to prevent backward compatibility issues.

```go get github.com/stacktitan/ldapauth```

## The go fmt Command

- Automatically formats your source code. (enforcing the use of proper line break, indentation and brace alignment).

```go fmt /path/to/your/package```

## The golint and go vet Commands

- golint reports style mistakes such as missing comments, variable naming that doesn't follow conventions, useless type specifications and more.
- golint is Standalone tool and will need to be installed.


```go get -u golang.org/x/lint/golint```

- go vet attempts to identify issues, some of which might be legitimate bugs, that a compiler might miss.

## Go Playground

- Is an execution enviornment hosted at [https://play.golang.org] that provides a web-based frontend for developers.

## Other Commands and Tools

- go test is used to run unit test and benchmarks, 
- go cover to check for test coverage
- go imports to fix import statements.


## Go Syntax

- [https://tour.golang.org]

### Data Types - Primitive Data Types 

- Primitive Data Types: bool, string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, byte, rune, float32, float64, complex64, complex128.

```var x = "Hello Wrold!" 
z := int(42)
```

### Data Types - Slices and Maps

- Slices are like arrays that you can dynamically resize and pass to functions more efficiently.
- Maps are associative arrays, unordeded lists of key/value pairs that allow you to efficiently and quickly look up values for a unique key.

``` var s = make([]string, 0)
var m = make([string]string)
s = append(s, "some string")
m["some key"] = "some value"
```

### Data Types - Pointers, Structs, and Interfaces 

- A pointer points to a particular area in memory and allows you to retrieve the value stored there.
- Use the & to retrieve the address in memory of some variable
- Use the * to dereference the address

``` var count = int(42)
ptr := &count
fmt.Println(*ptr)
*ptr = 100
fmt.Println(count)
```

- Use the struct type to define new data types 
- Use the new keyword to implement a new struct Person

``` 
type Person struct {
    Name string 
    Age int
}

func (*p Person) SayHello() {
    fmt.Println("Hello, ", p.Name")
} 

func main() {
    var guy = new (Person)
    guy.Name = "Dave"
    guy.SayHello()
}
```

- Go's interface type is a blueprint or a contract, that defines an expected set of actions that any concrete implementation must fulfill in order to be considered a type of that interface. 

```
type Friend interface {
    SayHello()
}
``` 

- The below takes the Friend interface as an input to the method. 

```
func Greet (f Friend) {
    f.SayHello()
}
```

- As Person type implements the Friend interface you can also pass it to the Greet method.

``` 
func main() {
    var guy = new(Person)
    guy.Name = "Dave"
    Greet(guy)
}
```

- Using interfaces and structs, you can define multiple types that you can pass to the same Greet() function, as long as these types implement the Friend interface.

```
type Dog struct {}
func (d *Dog) SayHello() {
    fmt.Println("Woof Woof!")
}

func main() {
    var guy = new(Person)
    guy.Name = "Dave"
    Greet(guy)
    var dog = new(Dog)
    Greet(dog)
}
```

## Control Structures