package main

import (
	"errors"
	"fmt"
)

// Struct
type Customer struct {
	Name, Address 	string
	Age           	int
	Married		bool
}

// Struct Method
func (customer Customer) sayHi(name string){
	fmt.Println("Hai, ", name, " My Name is ", customer.Name)
}

func (a Customer) sayHuuu(){
	fmt.Println("Huuuuu from ", a.Name)
}

// Interface
type HasName interface {
	GetName() string
}

func SayHello(hasName HasName){
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

// Interface Kosong
func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

// Nil
func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name":name,
		}
	}
}

// error Interface
func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

// Type Assertions
func random() interface{} {
	return "Egi"
}

func main() {
	// Struct
	var egi Customer
	egi.Name = "Egi"
	egi.Address = "Subang"
	egi.Age = 25

	fmt.Println(egi.Name)
	fmt.Println(egi.Address)
	fmt.Println(egi.Age)
	fmt.Println("---")
	
	// gilang := Customer{
	// 	Name: "Gilang",
	// 	Address: "Bandung",
	// 	Age: 24,
	// }
	// fmt.Println(gilang)
	
	// budi := Customer{"Budi", "Jakarta", 30, true}
	// fmt.Println(budi)

	// Struct Method
	egi.sayHi("Egi")
	egi.sayHuuu()
	fmt.Println("---")
	
	// Interface
	var nugraha Person
	nugraha.Name = "Nugraha"
	SayHello(nugraha)
	
	cat := Animal {
		Name: "Puss",
	}
	SayHello(cat)
	fmt.Println("---")
	
	// Interface Kosong
	var data interface{} = Ups(3)
	fmt.Println(data)
	fmt.Println("---")
	
	// Nil
	var personNil map[string]string = NewMap("Egi")
	if personNil == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(personNil)
	}
	fmt.Println("---")
	
	// error Interface
	hasil, err := Pembagi(100, 10)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
	fmt.Println("---")

	// Type Assertions
	var result interface{} = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown type")
	}
}