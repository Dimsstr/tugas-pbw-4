package main

import "fmt"

type Customer struct {
	Name    string
	Address string
	Age     int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	var eko Customer
	eko.Name = "Eko"
	eko.Address = "Indonesia"
	eko.Age = 25
	fmt.Println(eko)
	fmt.Println(eko.Name)
	fmt.Println(eko.Address)
	fmt.Println(eko.Age)
	joko := Customer{
		Name:    "Joko", // Memperbaiki pemisah antar field dengan koma ","
		Address: "Indonesia",
		Age:     30, // Memperbaiki pemisah antar field dengan koma ","
	}
	fmt.Println(joko)

	budi := Customer{Name: "Budi", Address: "Indonesia", Age: 30} // Menggunakan curly braces {} untuk membuat instance struct
	fmt.Println(budi)

	budi.sayHello("Agus")
	eko.sayHello("Agus")
	joko.sayHello("Agus")
}
