package main

/*import (
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
}

func (p *Person) CatchACold() {
	p.Fire(p.Name)
}

func NewPerson(name string) *Person {
	return &Person{
		Observable: Observable{new(list.List)},
		Name:       name,
	}
}

type DoctorService struct{}

func (d *DoctorService) Notify(data interface{}) {
	fmt.Printf("A Doctor has been called for %s \n", data.(string))
}

func main() {
	p := NewPerson("Guillermo")
	d := &DoctorService{}
	d2 := &DoctorService{}
	p.Subscribe(d)
	p.Subscribe(d2)
	p.CatchACold()
}
*/
