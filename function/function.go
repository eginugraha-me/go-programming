package main

import "fmt"

// Function Basic
func sayHello() {
	fmt.Println("Hello World")
}

// Function Parameter
func sayHelloTo(firstname string, lastname string) {
	fmt.Println("Hello ", firstname, lastname)
}

// Function Return Value
func getHello(name string) string {
	if name == "" {
		return "Hello Bro"
	} else {
		return "Hello " + name
	}
}

// Returning Multiple Values
func getFullname() (string, string, string){
	return "Egi", "Ar", "Nugraha"
}

// Named Return Values
func getFullname2() (firstNa string, middleNa string, lastNa string) {
	firstNa = "Egi"
	middleNa = "Ar"
	lastNa = "Nugraha"
	return 
}

// Variadic Function
func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers{
		total += value
	}
	return total
}

// Function as Value
func getGoodBye(name string) string {
	return "Good Bye, " + name
}

// Function as Parameter
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello, ", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

// Anonymous Function
type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked ", name)
	} else {
		fmt.Println("Welcome ", name)
	}
}

// func blacklistAdmin(name string) bool {
// 	return name == "admin"
// }

// func blacklistRoot(name string) bool {
// 	return name == "root"
// }

// Recursive Function
func factorialLoop(value int) int {
	result := 1
	for i:= value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value - 1)
	}
}

// Defer, Panic & Recover
func logging(){
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int){
	defer logging()
	fmt.Println("Run application")
	result := 10 / value
	fmt.Println("Result ", result)
}

func endApp(){
	message := recover()
	if message != nil {
		fmt.Println("Error dengan pesan: ", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool){
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	// Function Basic
	sayHello()
	fmt.Println("---")

	// Function Parameter
	firstName := "Egi"
	sayHelloTo(firstName, "Nugraha")
	sayHelloTo("Budi", "Budiman")
	fmt.Println("---")

	// Function Return Value
	result := getHello("Egi")
	fmt.Println(result)
	fmt.Println(getHello(""))
	fmt.Println("---")

	// Returning Multiple Values
	firstN, _, lastN := getFullname()
	fmt.Println(firstN)
	// fmt.Println(middleN)
	fmt.Println(lastN)
	fmt.Println("---")

	// Named Return Values
	a, b, c := getFullname2()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println("---")

	// Variadic Function
	totalS := sumAll(10, 90, 30, 50, 40)
	fmt.Println(totalS)
	
	slice := []int{10, 20, 30, 40, 50}
	totalS = sumAll(slice...)
	fmt.Println(totalS)
	fmt.Println("---")

	// Function as Value
	sayGoodBye := getGoodBye
	resultS := sayGoodBye("Egi")
	fmt.Println(resultS)
	fmt.Println(getGoodBye("Egi"))
	fmt.Println("---")
	
	// Function as Parameter
	sayHelloWithFilter("Egi", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)
	fmt.Println("---")
	
	// Anonymous Function
	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("egi", blacklist)
	registerUser("root", func(name string) bool {
		return name == "root"
	})
	registerUser("egi", func(name string) bool {
		return name == "root"
	})
	fmt.Println("---")
	
	// Recursive Function
	loop := factorialLoop(10)
	fmt.Println(loop)
	//fmt.Println(5 * 4 * 3 * 2 * 1)

	recursive := factorialRecursive(10)
	fmt.Println(recursive)
	fmt.Println("---")
	
	// Closure
	nameC := "Egi"
	counter := 0
	
	increment := func() {
		nameC = "Budi"
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(nameC)
	fmt.Println("---")
	
	// Defer, Panic & Recover
	runApplication(10)
	runApp(true)
	fmt.Println("Egi")
	fmt.Println("---")
}