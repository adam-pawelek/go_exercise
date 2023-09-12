package main

import "fmt"

func factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result = result * i
	}
	return result
}

func main() {

	var n int

	fmt.Print("Enter factorial number: \n")

	fmt.Scanf("%d", &n)

	fmt.Print(factorial(n))
	fmt.Print("\n")

}
