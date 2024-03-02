package main

import "fmt"

func main() {
	fmt.Println("Hi there!")
}

// Questions

// 1). How do we run the code in the project?
// In the terminal < go run main.go

// 2). What does 'import "fmt" mean?
// used to give our package access to some code written in another package

// 3). What does 'package main' mean?
// A collection of common source code files

// 4). Whats the 'func' thing?
// declares a function, sets the name of the function, list the arguments and call the functionn to run the code

// 5). How is the main.go file organized?
// package main - import "fmt" - func (etc.)

// --------------------------------------------------

// Go CLI commands

// go build - compiles a bunch of Go source code files

// go run - compiles and executes one or two files

// go fmt - Formats all the code in each file in the current directory

// go install - compiles and installs a package

// go get  -downloads the raw source code of someone elses package

// go test - runs any tests associated with the currrent project

// --------------------------------------------------

// Go Packages

// a package is a collection of common source code files 

// executable and reusable 

// main is used to create an executable package