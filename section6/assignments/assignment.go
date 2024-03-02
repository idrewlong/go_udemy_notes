// Write a program that creates two custom struct types called 'triangle' and 'square'
// The 'square' type should be struct with a field called 'sideLength' of type float64
// The triangle type should be a struct with a field called 'height' of type float64 and a field of type base of type float64
// Both types should have a function called 'getArea' that returns the calculated area of the square or triangle
// Area of a triangle = 0.5 * base * height
// Area of square = sideLength * sideLength
// Add a shape interface that defines a function called 'printArea' This function should calculate the area of a given shape and print it out to the terminal Design the interface so that the 'printArea' function can be called with either a triangle or square

// 1st Attempt
package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base float64
}

type square struct {
	sideLength float64
}

func main () {
	tr := triangle{}
	sq := square{}
}

// didn't print here 

func printArea(s shape) {
	fmt.Println(s.getArea())
}
// messed up here 

func (tr triangle) getArea() float64 {
	return (0.5 * tr.base * tr.height)
}
// didn't need the ()

func (sq square) getArea() float64 {
	return (sq.sideLength * sq.sideLength)
}
// didn't need the ()

// 2nd Attempt
package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base float64
}

type square struct {
	sideLength float64
}

func main() {
	t := triangle{ base: 10, height: 10}
	s := square{sideLength: 10}

	printArea(t)
	printArea(s)
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}