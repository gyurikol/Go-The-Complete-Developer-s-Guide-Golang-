package main

import "fmt"

/*
[0001] [ person{firstName: "Jim"..} ]
  ^                ^
address          value

Turn [address] into {value}   with *address     e.g. (*pointerToPerson).firstname = newFirstName
Turn {value}   into [address] with &value       e.g. jimPointer := &jim
*/

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstname string
	lastname  string
	contactInfo
}

func main() {
	jim := person{
		firstname: "Jim",
		lastname:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	//jimPointer := &jim // get pointer
	// jimPointer.updateName("Jimmy")
	jim.updateName("Jimmy")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstname = newFirstName
	// ^ get pointer address value
}

/*
	VALUE TYPES | REFERENCE TYPES
	------------+----------------
	    int     |    slices
	   float    |     maps
	  string    |   channels
	   bool     |   pointers
	  structs   |   functions
*/
