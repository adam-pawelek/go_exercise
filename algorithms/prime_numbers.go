package main

import (
	"fmt"
	"math"
)

func isPrimeNumber(number int) bool {

	if number <= 3 {
		return true
	}
	for i := 2; i < int(math.Sqrt(float64(number)))+1; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func main() {

	fmt.Print("Enter number of numbers \n")
	var numberOfNumbers int

	fmt.Scanf("%d", &numberOfNumbers)

	fmt.Print("Enter numbers\n")

	for i := 0; i < numberOfNumbers; i++ {
		var number int
		fmt.Scanf("%d", &number)
		fmt.Print(isPrimeNumber(number))
		fmt.Print("\n")
	}

}
