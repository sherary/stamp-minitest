package main

import (
	"fmt"
	"math"
	"strings"
)

func isPrime(num int) bool {
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}

	return num > 1
}

func numbersWithNoPrime(start, end int) []int {
	var nonPrime []int
	for i := start; i >= end; i-- {
		if !isPrime(i) {
			nonPrime = append(nonPrime, i)
		}
	}

	return nonPrime
}

func fooBar(numbers []int) string {
	var result []string

	for _, number := range numbers {
		switch {
		case number%3 == 0 && number%5 == 0:
			result = append(result, "FooBar")
		case number%3 == 0:
			result = append(result, "Foo")
		case number%5 == 0:
			result = append(result, "Bar")
		default:
			result = append(result, fmt.Sprintf("%d", number))
		}
	}

	return strings.Join(result, ", ")
}

func main() {
	prime := numbersWithNoPrime(100, 1)
	str := fooBar(prime)
	fmt.Println(str)
}
