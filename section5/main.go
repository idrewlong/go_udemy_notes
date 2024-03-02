package main

import "fmt"

func main(){
	colors := map[string]string{
		"red": "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c{
		fmt.Println("Hex code for", color, "is", hex)
	}
}




// func main() {
// 	colors :=make(map[string]string)

// 	colors["white"] = "#ffffff"

// 	delete(colors, "white")

// 	fmt.Println(colors)
// }



// func main() {
// 	var colors map[string]string
// }

// func main() {
// 	colors := map[string]string{
// 		"red": "#ff0000",
// 		"green": "#4bf745",
// 	}
// 	fmt.Println(colors)
// }