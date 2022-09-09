package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	d := gob.NewDecoder(&b)
	result := Person{}
	_ = d.Decode(&result)
	return &result
}

func (a *Address) DeepCopy() *Address {
	return &Address{
		StreetAddress: a.StreetAddress,
		City:          a.City,
		Country:       a.Country,
	}
}

func main() {
	name1 := Person{
		Name: "Guillermo",
		Address: &Address{
			StreetAddress: "London",
			City:          "Manchester",
			Country:       "UK",
		},
		Friends: []string{"jane", "does", "page"},
	}

	name2 := name1.DeepCopy()

	name2.Name = "Jane"

	name2.Address.StreetAddress = "London 23 23"
	name2.Friends = []string{"Guillermo"}
	fmt.Println(name1, name1.Address)
	fmt.Println(name2, name2.Address)
}
