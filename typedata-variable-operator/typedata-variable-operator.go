package main

import "fmt"

func main(){
	// Number
	fmt.Println("Satu = ", 1)
	fmt.Println("Dua = ", 2)
	fmt.Println("Tiga Koma Lima = ", 3.5)
	fmt.Println("--")

	// Bool
	fmt.Println("Benar = ", true)
	fmt.Println("Salah = ", false)
	fmt.Println("--")
	
	// String
	fmt.Println("Egi")
	fmt.Println("Egi Nugraha")
	fmt.Println(len("Egi"))
	fmt.Println("Egi Nugraha"[0])
	fmt.Println("--")

	// Variable
	var name string
	name = "Egi"
	fmt.Println(name)
	name = "Egi Nugraha"
	fmt.Println(name)

	var friendName = "Gilang"
	fmt.Println(friendName)
	var age = 25
	fmt.Println(age)
	
	country := "Indonesia"
	fmt.Println(country)

	var (
		firstName = "Egi"
		lastName = "Nugraha"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println("--")

	// Constant
	const (
		firstN string 	= "Egi"
		lastN			= "Nugraha"
		valueN		= 1000
	)
	fmt.Println(firstN)
	fmt.Println(lastN)
	fmt.Println(valueN)
	fmt.Println("--")

	// Type Data Conversion
	var nilai32 int32 = 129
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var nameC = "Egi"
	var e byte = nameC[0]
	var eString string = string(e)

	fmt.Println(nameC)
	fmt.Println(eString)
	fmt.Println("--")

	// Type Declarations
	type NoKTP string
	type Married bool

	var noKtpEgi NoKTP = "291873981723912213"
	var marriedStatus Married = false
	fmt.Println(noKtpEgi)
	fmt.Println(marriedStatus)
	fmt.Println("--")

	// Math Operator
	var result = 10 + 10
	fmt.Println(result)

	var a = 10
	var b = 10
	var c = a * b
	fmt.Println(c)

	var i = 10
	i += 10
	fmt.Println(i)

	i++
	fmt.Println(i)

	var negative = -100
	var positive = 100
	fmt.Println(negative)
	fmt.Println(positive)
	fmt.Println("--")

	// Comparation Operator
	var name1 = "Egi"
	var name2 = "Gilang"

	var resultN bool = name1 > name2
	fmt.Println(resultN)

	var value1 = 100
	var value2 = 200
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
	fmt.Println("--")

	// Bool Operator
	var ujian = 80
	var absensi = 75

	var lulusUjian = ujian >= 80
	var lulusAbsensi = absensi >= 80
	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)

	var lulus = lulusAbsensi && lulusUjian
	fmt.Println(lulus)
	fmt.Println(ujian >= 80 && absensi >= 80)
}