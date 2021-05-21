package main

import (
	"html/template"
	"os"
)

type User struct {
	Name       string
	Dog        map[string]string
	Age        int
	Hobbies    []string
	LikeGolang bool
}

func main() {
	t, err := template.ParseFiles("exp/hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := User{
		Name: "John Smith",
		Age:  25,
		Dog: map[string]string{
			"Name":  "Snoopy",
			"Breed": "Beagle",
		},
		Hobbies:    []string{"running", "cooking"},
		LikeGolang: true,
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
