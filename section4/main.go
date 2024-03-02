package main

import "fmt"

type contactInfo struct {
	email string
	zip int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName: "Jones",
		contactInfo: contactInfo{
			email: "jimjones@gmail.com",
			zip: 37288,
		},
	}
	
	jim.updateName("Jimmy")
	jim.print()
}

func (pointertoPerson *person) updateName(newFirstName string){
	(*pointertoPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)

}

// Type #1 for Declaring Structs
// func main(){
// 	alex := person{"Alex", "Anderson"}
// 	fmt.Println(alex)
// }

// Type #2 for Declaring Structs
// func main(){
// 	alex := person{firstName: "Alex", lastName: "Anderson"}
// 	fmt.Println(alex)
// }

// Type #3 for Declaring Structs
// func main(){
// 	var alex person
// 	alex.firstName = "Alex"
// 	alex.lastName = "Anderson"
// 	fmt.Println(alex)
// 	fmt.Printf("%+v", alex)
// }



