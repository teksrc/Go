package main

import (
	"fmt"
	"html/template"
	"os"
)

type User struct {
	Name string
	Bio string
	Age int
	// Meta UserMeta
}

type UserMeta struct {
	Visits int
}

func main() {
	fmt.Println("Experimental: ")
	t, err:= template.ParseFiles("hello.html")
	if err != nil {
    panic(err)
	}
	user := User {
		Name: "Morpheus",
		Bio: "The finder of the One",
		Age: 36,
		// Meta: UserMeta {
		// 	Visits: 1,
		// },
	}
	
	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
