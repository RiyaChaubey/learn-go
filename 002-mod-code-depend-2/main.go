package main

/* How import of puppy module was possible?
1. create a github repo puppy and add your puppy.go file and package
3. to initialiaze puppy repo as a go module run - go init github.com/RiyaChaubey/puppy
2. push that
3. come in this repo (where you want to import puppy)
4. make you folder/file structure, then run - go get github.com/RiyaChaubey/puppy@latest
5. go mod tidy (this will import the dependency module in this module's go.mod file )
6. Rest of the steps are this file
7. You can your program by - go run main.go (be in the folder 001-mod-code-depend)
*/

/*
When a module which is being imported gets outdated, we need to run
go get github.com/RiyaChaubey/puppy@commit again, to get the latest code from module.
*/

/*
This shows double dependecy of a module
learn-go-002 <- puppy <- dog
since dog in imported in puppy and puppy is imported in learn-go
you can use "exported" (first letter captial var/func) things in learn-go
of both puppy and dog just by importing puppy
as there is a nested dependency here.
*/

import (
	"fmt"

	"github.com/RiyaChaubey/puppy"
)

func main() {
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	fmt.Println(s1)
	fmt.Println(s2)

	s3 := puppy.BigBark()
	s4 := puppy.BigBarks()

	fmt.Println(s3)
	fmt.Println(s4)
}
