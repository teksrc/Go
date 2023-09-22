package main

import (
	"fmt"
)

type User struct {
	Name string
	Bio string
	Age int
	Meta UserMeta
}

type UserMeta struct {
	Visits int
}

func main() {
	fmt.Println("Experimental: ")
	// user := User {
	// 	Name: "Morpheus",
	// 	Bio: "The finder of the One",
	// 	Age: 36,
	// 	Meta: UserMeta {
	// 		Visits: 1,
	// 	},
	// }
	
}
