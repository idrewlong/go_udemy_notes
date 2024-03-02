// Create a program that reads the contents of a text file then prints its contents to a terminal
// The file to open should be provided as an argument to the program when it is executed at the terminal. For example 'go run asssignment2.go myfile.txt' should open up the myfile.txt file
// to read in the arguments provided to a program, you can reference the variable 'os.Args' which is a slice of type string
// To open a file, check out the documentation for the 'Open' function in the 'os' package
// What interfaces does the 'File' type implement?
// If the 'File' type implements the 'Reader' interface, you might be able to reuse that io.Copy fucntion

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil{
		fmt.Println("Error", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}

// ./assignment2 assignment2.go