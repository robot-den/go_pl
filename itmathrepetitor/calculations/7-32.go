package main

import (
	"fmt"
	"math"
)

func main() {
	ex7()
	ex8(2, 3)
	ex9(-2)
	ex10(1.7)
	ex10(3)
	ex11(-2.34)
	ex12(3.6)
	ex13(0.1, 0.2, 1.0)
	//ex14()
	//ex15()
	//ex16()
	ex17()
}

func ex7() {
	fmt.Println("Ex 7:")
	fmt.Println(1/2.0 + 1/4.0)
}

func ex8(a, b float64) {
	fmt.Println("\nEx 8:")
	fmt.Println((a+4*b)*(a-3*b) + a*a)
}

func ex9(x float64) {
	fmt.Println("\nEx 9:")
	fmt.Println(math.Abs(x) + math.Pow(x, 5))
}

func ex10(x float64) {
	fmt.Println("\nEx 10:")
	fmt.Println(math.Pow(x+1, 2) + 3*(x+1))
}

func ex11(x float64) {
	fmt.Println("\nEx 11:")
	a := ((math.Abs(x-5) - math.Sin(x)) / 3.0) + math.Sqrt(x*x+2014*math.Cos(2*x-3))
	fmt.Println(a)
}

func ex12(x float64) {
	fmt.Println("\nEx 12:")
	a := math.Pow(math.E, x-2) + math.Abs(math.Sin(x)) - math.Pow(x, 4)*math.Cos(1.0/x)
	fmt.Println(a)
}

func ex13(a, b, x float64) {
	fmt.Println("\nEx 13:")

	p3 := (x * x) + b - (((b * b) * math.Pow(math.Sin(x+a), 3)) / x)
	fmt.Println(math.Pow(p3, 1.0/5.0))
}

func ex14() {
	fmt.Println("\nEx 14:")
	fmt.Println("set a and b in a format 'a,b'")
	//reader := bufio.NewReader(os.Stdin)
	//input, err := reader.ReadString('\n')
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//rawValues := strings.Split(strings.TrimSuffix(input, "\n"), ",")
	//if len(rawValues) != 2 {
	//	fmt.Println("Wrong input format")
	//	return
	//}
	//a, err := strconv.Atoi(rawValues[0])
	//b, err := strconv.Atoi(rawValues[1])
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	var a, b int
	if _, err := fmt.Scanf("%d,%d\n", &a, &b); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("sum: %d\n", a+b)
	fmt.Printf("mul: %d\n", a*b)
}

func ex15() {
	fmt.Println("\nEx 15:")
	fmt.Println("set x")
	var x float64
	if _, err := fmt.Scanf("%f\n", &x); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("x^2: %d\n", int(math.Pow(x, 2)))
	fmt.Printf("x^3: %d\n", int(math.Pow(x, 3)))
}

func ex16() {
	fmt.Println("\nEx 16:")
	fmt.Println("set a, b and c in a format 'a,b,c'")
	var a, b, c int
	if _, err := fmt.Scanf("%d,%d,%d\n", &a, &b, &c); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("res:", (a*2)+(b-3)+(c*c))
}

func ex17() {
	fmt.Println("\nEx 17:")
	fmt.Println("set a, b and c in a format 'a,b,c'")
	var a, b, c int
	if _, err := fmt.Scanf("%d,%d,%d\n", &a, &b, &c); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("mean:", (a+b+c)/3.0)
	fmt.Println("calc:", ((a+b)*2)-(c*3))
}
