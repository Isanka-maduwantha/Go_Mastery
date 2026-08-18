package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)
// 22. Configuration Builder: Create a struct-like constant set for app configuration
type AppConfig struct {
	port int
	Env string
	DB string
}
func main() {
	// variableMutation()
	// unScope()
	// fmt.Println(username)
	// varTypes()
	// constants()
	// Iota()
	// birthdayProgram()
	// birthdayProrefined()
	// varInfoPrint()
	// constantCalculator()
	// typeExplorer()
	// variableSwapper()
	// scoopTester()
	// unitConverter()

	config := AppConfig{
		Env: "Hello",
		port: 3000,
		DB : "Ehllo Again",
	}
	fmt.Println("Server Running on PORT: ",config.port)
}
func warmUp() {
	// WarmUp

	// Declare 5 variables using var keyword with different types (int, string, bool, float64, byte
	var text string = "Hello"
	var num int = 12
	var isTrue bool = true
	var price float64 = 3.14
	var letter byte = 'A'
	fmt.Println(text, num, isTrue, price, letter)
	// Declare 5 variables using short declaration := operator
	text2 := "hello"
	num2 := 123
	isThere := false
	price2 := 341.70
	byte8 := byte('A')
	fmt.Printf(text2, num2, isThere, price2, byte8)
	var x int = 5
	var y = 10
	z := 12

	fmt.Print(x, y, z)
}

/*
 Core Concepts
*/

// 6. Create a program that declares a variable, reassigns it, and prints before/after values
func variableMutation() {
	x := 10
	fmt.Println("Before: ", x)
	x = 12
	fmt.Println("After:  ", x)

}

var username = "Isanka"

// 7. Understand scope: declare variables at package and function level, observe access
func unScope() {
	username := username + " Kumara"
	fmt.Println(username)
}

// 8. Create variables with explicit types: var name string = "Go"
// 9. Create variables with implicit types: name := "Go" and check type
// 10. Initialize multiple variables in one statement: var x, y, z int = 1, 2, 3
func varTypes() {
	var name string = "GO"
	// var x,y,z = 10,"Hello",true
	var m, l, n float64

	fmt.Println(m, l, n)
	uname := "GO"
	fmt.Println(name, uname)
}

// 12. Create named constants: const Pi = 3.14159
// 13. Create grouped constants with const ()
func constants() {
	const PI = 3.14
	const HI, TIME, YEAR = "H!", "12:30", 2026

	// HI = "HI" // cannot assign to HI (neither addressable nor a map index expression
	fmt.Printf(HI)
}

// 14. Understand iota: create enum-like constants (0, 1, 2, 4, 8...)
func Iota() {
	const (
		_ = iota
		item1
		item2
		_
		item3
	)
	fmt.Println(item1, item2, item3)
}

// 15. Create a "birthday" program: declare DOB as string, calculate age using constants
func birthdayProgram() {
	var DOB string
	fmt.Print("Enter Your Birthday(YYYY-MM-DD): ")
	fmt.Scanln(&DOB)

	year, month, day := time.Now().Date()
	result := strings.Split(DOB, "-")
	YEAR, _ := strconv.Atoi(result[0])
	MONTH, _ := strconv.Atoi(result[1])
	DAY, _ := strconv.Atoi(result[2])
	fmt.Printf("TODAY: %d %d %d \n", year, month, day)
	fmt.Printf("BIRTHDAY : %d %d %d\n", YEAR, MONTH, DAY)
	AGE := year - YEAR
	fmt.Println("Your Age: \n", AGE, " Years")

}
func birthdayProrefined() {
	const DOBStr = "2003-06-20"
	const LayoutStr = "2006-01-02"
	dob, err := time.Parse(LayoutStr, DOBStr)
	if err != nil {
		fmt.Println("Error Parsing Date: ", err)
		return
	}
	now := time.Now()

	age := now.Year() - dob.Year()

	fmt.Println("Date of Birth: ", DOBStr)
	fmt.Println("Current Age: ", age)
}

/* Exersices */

// 16. Variable Info Printer: Create a program that takes 3 variables and prints their type, value, and memory address

func varInfoPrint() {
	var val1 string
	var val2 int
	var val3 bool

	// TakeString int and bool values
	fmt.Print("Enter Name : ")
	fmt.Scanln(&val1)

	fmt.Print("Enter Number : ")
	_, err := fmt.Scanln(&val2)
	// check if the input is in valid format
	if err != nil {
		fmt.Println("Enter valid Number: ", err)
		return
	}

	fmt.Print("Are You A Student(true/false): ")
	fmt.Scanln(&val3)

	// Print Variable value,Type,Memory Address
	fmt.Printf("Name: %s\nType: %T\n Pointer: %p\n\n", val1, val1, &val1)

	fmt.Printf("Number: %d\nType: %T\n Pointer: %p\n\n", val2, val2, &val2)

	fmt.Printf("Is Student: %v\nType: %T\n Pointer: %p\n\n", val3, val3, &val3)

}

// 17. Constant Calculator: Define constants for geometric shapes and calculate properties
func constantCalculator() {
	const (
		PI = 3.14
		r  = 11
	)

	circumference := 2 * PI * r
	Area := 2 * PI * (r * r)

	fmt.Printf("Circle: \n  Circumference: %0.2f\n  Area: %0.2f\n", circumference, Area)
}

// 18. Type Explorer: Write a program that demonstrates all basic types and their zero-values
func typeExplorer() {
	var num int
	var name string
	var isValid bool
	var price float64
	var base8 byte
	fmt.Println("Type and Their Zero value")
	fmt.Printf("Type : %T\nZero value: %d\n\n", num, num)
	fmt.Printf("Type : %T\nZero value: %s\n\n", name, name)
	fmt.Printf("Type : %T\nZero value: %v\n\n", isValid, isValid)
	fmt.Printf("Type : %T\nZero value: %f\n\n", price, price)
	fmt.Printf("Type : %T\nZero value: %v\n\n", base8, base8)
}

// 19. Variable Swapper: Swap two variables without using a temp variable (use multiple assignment)
func variableSwapper() {
	a := 10
	b := 15
	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)
}

// 20. Scope Tester: Create nested functions and demonstrate variable shadowing
func scoopTester() {
	var name string = "SShehan"

	sayHello := func() {
		// Variable Shadowing
		var name string = "Isanka"
		fmt.Println("Hi ", name)
	}
	sayHello()
	fmt.Println(name)
}

// 21. Units Converter: Create constants for unit conversion (e.g., 1 mile = 1.60934 km)
func unitConverter() {
	const (
		mile    =  1.60934
	)
	fmt.Println("I Want to Convert:       ")
	fmt.Println("1.Milimeter to Centimeter")
	fmt.Println("2.Centimeter to Meter   ")
	fmt.Println("3.Kilometer to meter     ")
	fmt.Print("4.Mile to Kilometer      \n\n")
	fmt.Print("Make Your Pick:")

	var val,x int
	_, err := fmt.Scanln(&val)
	if err != nil {
		fmt.Println("Error Occured", err)
		return
	}
	fmt.Println("You Selected", val)
	fmt.Print(x)

}
