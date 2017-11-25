package main

import (
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int
}

// basically, DoubleZero inherits from Person,
// it has access to all fields of Person
type DoubleZero struct {
	Person        // anonymous type
	First         string
	LicenseToKill bool
}

func main() {
	p1 := DoubleZero{
		Person: Person{
			First: "James",
			Last:  "Bond",
			Age:   20,
		},
		First:         "Double Zero Seven",
		LicenseToKill: true,
	}

	p2 := DoubleZero{
		Person: Person{
			First: "Miss",
			Last:  "Moneypenny",
			Age:   18,
		},
		First:         "If looks could kill",
		LicenseToKill: false,
	}

	//fields and methods of thr inner-type are promoted to the outer type
	fmt.Println(p1.First, p1.Person.First)
	fmt.Println(p2.First, p2.Person.First)
	// p1 := person{"James", "Bond", 20}
	// p2 := person{"Miss", "Moneypenny", 18}

	// fmt.Println(p1.first, p1.last, p1.age)
	// fmt.Println(p2.first, p2.last, p2.age)
}

// Go is object-oriented
// (1) encapsulation
// state (fields)
// behaviour (methods)
// exported/unexported
//
// (2) Reusability
// inheritance (embedded types)
//
// (3) Polymorphism
// interfaces

// (4) Overriding
// "promotion"
