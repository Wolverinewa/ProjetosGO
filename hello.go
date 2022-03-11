package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var c, d string = "Willian", "Danilo"
	fmt.Printf("hello, world\n")
	fmt.Printf("hello, %v \n", c)
	fmt.Printf("hello, %v \n", d)
	fmt.Printf("hello, %v e %v \n", d, c)

	type Carro struct {
		Name  string
		Year  int
		Color string
	}

	car := Carro{"BMW 320i", 2021, "Preto"}
	result, _ := json.Marshal(car) // dados estão em bytes
	fmt.Println(string(result))    //transformando dados em string

	var car2 Carro
	json := []byte(`{"Name":"BMW 320","Year":2021,"Color":"Black"}`)

	json.Unmarshal(json, car2)
	fmt.Println(car2.Name) // BMW 320
}
