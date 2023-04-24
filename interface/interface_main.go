package main

import (
	"fmt"
	"math"
)

type Stringer interface {
	String() string
}

type Sizer interface {
	Area() float64
}

type Shaper interface {
	Sizer
	Stringer
}

type Article struct {
	Title  string
	Author string
}

type Book struct {
	Title  string
	Author string
	Pages  int
}

type Circle struct {
	Radius float64
}

type Square struct {
	Width  float64
	Height float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (s Square) Area() float64 {
	return s.Width * s.Height
}

func (a Article) String() string {
	// fmt.Println("Title: ", a.Title, " Author: ", a.Author)
	return fmt.Sprintf("Title: %s, Author: %s", a.Title, a.Author)
}

func (b Book) String() string {
	return fmt.Sprintf("The %q book was written by %s.", b.Title, b.Author)
}

func (s Square) String() string {
	return fmt.Sprintf("Square {Width: %.2f, Height: %.2f}", s.Width, s.Height)
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle {Radius: %.2f}", c.Radius)
}

func main() {
	a := Article{
		Title:  "Understanding Interfaces in Go",
		Author: "Sammy Shark",
	}

	// fmt.Println(a.string())
	Print(a)

	b := Book{
		Title:  "All About Go",
		Author: "Jenny Dolphin",
		Pages:  25,
	}
	Print(b)

	c := Circle{Radius: 10}
	PrintArea(c)
	s := Square{Height: 10, Width: 5}
	PrintArea(s)

	l := Less(c, s)
	fmt.Printf("%+v is the smallest\n", l)
}

func Print(s Stringer) {
	fmt.Println(s.String())
}

func Less(s1, s2 Sizer) Sizer {
	if s1.Area() < s2.Area() {
		return s1
	}
	return s2
}

func PrintArea(s Shaper) {
	fmt.Printf("area of %s is %.2f\n", s.String(), s.Area())
}
