package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/Isanka-maduwantha/Go_Mastery/Level-02/handler"
)

func main() {
	handler.HandleHello()
	// WarmUp
	// performArithmetic()

	// Core concepts
	// comparisonTruthTable()
	// logicalExpressions()
	// shortCircuit()
	// compareTypes()
	// withinTolerance()
	// reverseString()

	// checkSubString()
	// converBetweenUnits()

	// c := CompoundIntrestVariables{
	// 	principal: 300000.00,
	// 	rate:      3.5,
	// 	nPeriods:  12,
	// 	time:      3,
	// }
	// calCompoundIntrest(c)
	// signChecker(-23+100)

	// Project Exersices

	// data :=CalculateData{
	// 	value1: 12,
	// 	operator: "-",
	// 	value2: 23,
	// }
	// fmt.Println(calculator(data))

	// calcBMI(80,1.62)

	// l  := Loan{
	// 	principal: 400000,
	// 	rate: 2.4,
	// 	years: 5,
	// }
	// l.monthlyPay =monthlyPayCalc(l)
	// fmt.Println(l)

	// gradeChecker(100)
	// stringAnalyzer("Heallo3423xIJa")
	// distanceCalc(10,10,67,99)
	// leapYearChecker(2008)
	// handler.ByeBye()
}

// Warmup

// 1.Perform all arithmetic operations: +, -, *, /, % on integers
func performArithmetic() {
	x, y := 15, 12
	m, n := 24.2, 12.5

	// Arithmetic Operations
	add := x + y
	sub := x - y
	multi := x * y
	divinr := x / y
	// 2.Perform arithmetic on float64 values
	div := (m / n)

	fmt.Print("1.Arithmetic Operations (x = 15, y = 12)\n")
	fmt.Println("Addition: ", add)
	fmt.Println("Subtraction: ", sub)
	fmt.Println("Multiplication: ", multi)
	fmt.Println("Division int: ", divinr)

	fmt.Printf("Division %T: %f ", div, div)

	// 3.Demonstrate operator precedence: 2 + 3 * 4 vs (2 + 3) * 4
	fmt.Print("\n\n2.Operator precedence\n")
	l := 2 + 3*4
	j := (2 + 3) * 4
	fmt.Println("2+3*4 =", l)
	fmt.Println("(2+3)*4 =", j)

	// 4.Show what happens with division by zero (what does it do?)
	// fmt.Println("Division By Zero ")
	// p,k := 12,0
	// r := p/k
	// fmt.Println("Result: ",r)

	// 5.Concatenate strings using + operator
	fname := "Isanka"
	lname := "Maduwantha"

	fmt.Println("FullName: ", fname+" "+lname)

}

// Core Concept Exercises
// 6. Create a comparison truth table: all 6 comparison operators on different types
func comparisonTruthTable() {
	x, y := 12.4, 11.8
	a := x == y
	b := x != y
	c := x > y
	d := x < y
	e := x >= y
	f := x <= y
	fmt.Println("Comparision X = ", x, " And ", " Y = ", y, "with 6 Comparison Operators:")
	fmt.Println("X equal to Y: ", a)
	fmt.Println("X not Equal to Y: ", b)
	fmt.Println("X is Greater than Y: ", c)
	fmt.Println("X is Less than Y: ", d)
	fmt.Println("X is Greater than or Equal to Y: ", e)
	fmt.Println("X is Less than or Equal to Y: ", f)

}

// 7. Write logical expressions using &&, ||, !
func logicalExpressions() {
	x, y, z := 13, 12, 24

	a := x > y && z > x || x != z
	b := y > z
	c := y*2 == z && (x*2)-2 == z
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}

// 8. Demonstrate short-circuit evaluation: false && expensive() shouldn't call expensive()
func shortCircuit() {
	expensive := func() bool {
		fmt.Println("It is Expensive")
		return true
	}
	var x float64
	fmt.Printf("Enter your Budget: ")
	fmt.Scanln(&x)
	isAble := x < 5000
	if isAble && expensive() {
		fmt.Println("You Cant Afford this")
	} else {
		fmt.Println("You Can Afford this")
	}
}

// 9. Compare different types: int vs float64, string vs string
func compareTypes() {

	// var x int = 10
	// var y float64 = 34.21
	var strX, strY = "Hello", "Mchn"

	// So you cannot compare Mismatched tyles
	// a := x != y
	b := strX != strY
	fmt.Println(b)

}

// 10. Create a "within tolerance" checker: abs(a - b) < epsilon
func withinTolerance() {
	tolerance := func(a, b, epsilon float64) bool {
		return math.Abs(a-b) < epsilon
	}
	a := 0.1
	c := 0.2
	d := a + c
	b := 0.3
	epsilon := 1e-9
	fmt.Printf("Is a %.17f  equal to %.17f : %v \n", d, b, tolerance(d, b, epsilon))

}

// 11. String manipulation: reverse a string using indexing
func reverseString() {
	s := "Hello Go"
	runes := []rune(s)

	fmt.Println(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	result := string(runes)
	fmt.Println(result)
}

// 12. String manipulation: check if a substring exists
func checkSubString() {
	s := "👋 Hello Go is A New Language to You"
	runes := []rune(s)
	// Exreacting subString
	sub1 := string(runes[0:7])
	sub2 := string(runes[8:])

	fmt.Println(sub1)
	fmt.Println(sub2)
	// Check if the Srting Exists
	chk := "Lang"
	if strings.Contains(string(runes), chk) {
		fmt.Println("Word " + chk + " is in The text")

	}

}

// 13. Use arithmetic to convert between units (temperature, distance)
func converBetweenUnits() {
	celToFer := func(Celc float64) float64 {
		return (Celc * 1.8) + 32
	}
	kmToM := func(Km float64) float64 {
		return Km * 1000
	}

	var (
		tempreture = 100.0
		distance   = 23.0
	)
	fmt.Println("Tempreture ", tempreture, "C ", " = ", celToFer(tempreture), "F")
	fmt.Println("Distance ", distance, "Km ", " = ", kmToM(distance), "m")
}

// 14. Calculate compound interest: A = P(1 + r/n)^(nt)
type CompoundIntrestVariables struct {
	principal float64
	rate      float64
	nPeriods  int
	time      int
}

func calCompoundIntrest(c CompoundIntrestVariables) {
	x := (c.nPeriods * c.time)
	r := c.rate / 100
	y := 1 + (r / float64(c.nPeriods))
	z := math.Pow(y, float64(x))
	A := c.principal * z

	fmt.Printf("Final Amount After %d  Years is : %.2f\n", c.time, A)
}

// 15. Create a "sign checker": return -1, 0, or 1 based on number sign
func signChecker(n float64) int {
	if n > 0 {
		fmt.Println(n, " is a Positive value")
		return 1
	} else if n < 0 {
		fmt.Println(n, " is a Negative value")
		return -1
	} else {
		fmt.Println(n, " is Equal to Zero(0)")
		return 0
	}
}

// Project Exercises (16-22)

// 16. Calculator Program: Create a basic +, -, *, / calculator
type CalculateData struct {
	value1   float64
	operator string
	value2   float64
}

// 17. BMI Calculator: height, weight → calculate and categorize BMI

func calculator(cdata CalculateData) float64 {
	x := cdata.value1
	y := cdata.value2
	switch cdata.operator {
	case "+":
		return x + y

	case "-":
		return x - y

	case "*":
		return x * y

	case "/":
		return x / y

	default:
		return 0
	}
}

func calcBMI(w float32, h float32) {
	BMI := w / (h * h)
	fmt.Println("BMI is ", BMI)

}

// 18. Loan Calculator: principal, rate, years → monthly payment

type Loan struct {
	principal  float64
	rate       float64
	years      int
	monthlyPay float64
}

func monthlyPayCalc(loan Loan) float64 {
	p := loan.principal
	r := loan.rate / 100
	y := loan.years

	z := (p * r) * float64(y)

	return z / float64(y*12)

}

// 19. Grade Checker: score → letter grade using comparison operators
func gradeChecker(m int) {
	g := "F"

	if m > 100 {
		fmt.Println("Invalid Score Value")
		return
	}
	if m >= 75 {
		g = "A"
	} else if m >= 60 {
		g = "B"
	} else if m >= 45 {
		g = "C"
	} else if m >= 35 {
		g = "S"
	}
	fmt.Println("Grade is : ", g)
}

// 20. String Analyzer: input string → count vowels, consonants, numbers
func stringAnalyzer(val string) {

	lowerVal := strings.ToLower(val)
	vov, cons, num := 0, 0, 0

	for _, i := range lowerVal {
		fmt.Println(i)
		if strings.ContainsRune("aeiou", i) {
			vov++
		} else if i >= 'a' && i <= 'z' {
			cons++
		} else if strings.ContainsRune("123456789", i) {
			num++
		}
		i++
	}
	fmt.Println("String :", val)
	fmt.Println("Vowels :", vov)
	fmt.Println("Consonents :", cons)
	fmt.Println("Numbers :", num)

}

// 21. Distance Calculator: x1, y1, x2, y2 → euclidean distance
func distanceCalc(x1 int, y1 int, x2 int, y2 int) {
	x := math.Pow(float64(x2-x1), 2)
	y := math.Pow(float64(y2-y1), 2)
	d := math.Sqrt(x + y)

	fmt.Printf("Euclidean Distance is : %.2fkm\n", d)

}

// 22. Leap Year Checker: year → is leap year? (complex logic)
func leapYearChecker(y int) {
	msg := ""
	if y%4 != 0 {
		msg = "not"
		
	}
	if y%100 == 0 && y%400 != 0 {
		msg = "not"
		
	}
	fmt.Println("Year ", y, " is", msg, "a Leap year")
}
