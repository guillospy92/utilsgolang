package main

type Observable interface {
	AddObserver(name string, observer Observer)
	RemoveObserver(name string)
	NotifyObserver()
}

type Observer interface {
	GetID() int
	Notify(data string)
}

type Person struct {
	observer map[string]Observer
	Name     string
}

func (p Person) AddObserver(name string, observer Observer) {
	p.observer[name] = observer
}

func (p Person) RemoveObserver(name string) {
	delete(p.observer, name)
}

func (p Person) NotifyObserver() {
	for _, v := range p.observer {
		v.Notify(p.Name)
	}
}

func NewPerson(name string) *Person {
	return &Person{make(map[string]Observer), name}
}
