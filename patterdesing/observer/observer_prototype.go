package main

import (
	"container/list"
	"fmt"
)

type Observer interface {
	Notify(data interface{})
}

type Observable struct {
	subscribe *list.List
}

func (o *Observable) Subscribe(x Observer) {
	o.subscribe.PushBack(x)
}

func (o *Observable) Unsubscribe(x Observer) {
	for z := o.subscribe.Front(); z != nil; z = z.Next() {
		if z.Value.(Observer) == x {
			o.subscribe.Remove(z)
		}
	}
}

func (o *Observable) Fire(data interface{}) {
	for z := o.subscribe.Front(); z != nil; z = z.Next() {
		z.Value.(Observer).Notify(data)
	}
}

type Person struct {
	Observable
	Name string
	age  int
}

func (p *Person) CatchACold() {
	p.Fire(p.Name)
}

func NewPerson(age int) *Person {
	return &Person{
		Observable: Observable{new(list.List)},
		age:        age,
	}
}

type PropertyChange struct {
	Name  string
	Value interface{}
}

func (p *Person) Age() int {
	return p.age
}

func (p *Person) SetAge(age int) {
	if age == p.age {
		return
	}

	p.age = age
	p.Fire(PropertyChange{"Age", p.age})
}

type TrafficManagement struct {
	o Observable
}

func (t *TrafficManagement) Notify(data interface{}) {
	if pc, ok := data.(PropertyChange); ok {
		if pc.Value.(int) > 16 {
			fmt.Println("Congrats, you can drive now")
			t.o.Unsubscribe(t)
		}
	}

}

func main() {
	p := NewPerson(15)
	t := &TrafficManagement{p.Observable}
	p.Subscribe(t)

	for i := 16; i <= 20; i++ {
		p.SetAge(i)
	}
}
