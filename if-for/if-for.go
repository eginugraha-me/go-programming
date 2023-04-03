package main

import "fmt"

func main() {
	// IF Expression
	var name = "Egi"

	if name == "Egi" {
		fmt.Println("Hello Egi")
	} else if name == "Joko" {
		fmt.Println("Hello Joko")
	} else if name == "Budi" {
		fmt.Println("Hello Budi")
	} else {
		fmt.Println("Hi, kenalan dong")
	}

	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}
	fmt.Println("--")

	// Switch Expression
	names := "Egi"

	switch names {
	case "Egi":
		fmt.Println("Hello Egi")
	case "Joko":
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Kenalan Dong")
	}

	switch lengths := len(names); lengths > 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Benar")
	} 

	lengthSw := len(names)
	switch {
	case lengthSw > 10:
		fmt.Println("Nama Terlalu Panjang")
	case lengthSw > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Sudah Benar")
	}
	fmt.Println("--")

	// For Loop

	// counter := 1
	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke ", counter)
	// 	counter++
	// }

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}

	slice := []string{"Egi", "Ar", "Nugraha", "Budi", "Joko"}
	for i := 0; i < len(slice); i++{
		fmt.Println(slice[i])
	}

	slice2 := []string{ "Budi", "Joko", "Egi", "Ar", "Nugraha"}
	for index, value := range slice2 {
		fmt.Println("Index", index, "=", value)
	}
	
	// Jika satu variable tidak terpakai, gunakan "_" underscore
	// for _, value := range slice2 {
	// 	fmt.Println(value)
	// }

	person := make(map[string]string)
	person["nama"] = "Egi"
	person["title"] = "Programmer"

	for key, values := range person {
		fmt.Println(key, "=", values)
	}
	fmt.Println("--")

	// Break & Continue
	for j := 0; j < 10; j++ {
		if j == 5 {
			break
		}
		fmt.Println("Perulangan ke ", j)
	}

	for k := 0; k < 10; k++ {
		if k % 2 == 0 {
			continue
		}
		fmt.Println("Perulangan ke ", k)
	}
}