package main

import (
	"fmt"
	"sync"
)

var once sync.Once

var instance *SingletonDatabase

type SingletonDatabase struct {
	capitals map[string]int
}

func (db *SingletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

func GetSingletonDatabase() *SingletonDatabase {
	once.Do(func() {
		fmt.Println(1, "Init app")
		capital := map[string]int{
			"w": 1,
			"a": 2,
			"b": 4,
			"d": 5,
		}
		instance = &SingletonDatabase{capital}
		fmt.Println(2, "in singleton")
	})

	return instance
}

func main() {
	db := GetSingletonDatabase()
	fmt.Println(db.GetPopulation("w"), "key database 1")

	db2 := GetSingletonDatabase()
	fmt.Println(db2.GetPopulation("d"), "key database 2")
}
