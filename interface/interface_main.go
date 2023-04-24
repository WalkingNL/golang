package main

import "fmt"

type Article struct {
	Title  string
	Author string
}

func (a Article) String() string {
	// fmt.Println("Title: ", a.Title, " Author: ", a.Author)
	return fmt.Sprintf("Title: %s, Author: %s", a.Title, a.Author)
}

func main() {
	a := Article{
		Title:  "Understanding Interfaces in Go",
		Author: "Sammy Shark",
	}

	// fmt.Println(a.string())
	print(a)
}

func print(s interface_test.Stringer) {
	fmt.Println(s.String())
}
