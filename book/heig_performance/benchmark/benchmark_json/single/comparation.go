package single

import (
	"encoding/json"
	"fmt"
)

type CarData struct {
	Cylinders int
	Brand     string
	Mpg       float64
}

func CarJSON() {
	honda := CarData{Cylinders: 4, Brand: "Toyota", Mpg: 42.6}
	var carDataJson []byte
	carDataJson, _ = json.Marshal(honda)
	fmt.Println(string(carDataJson))
}
