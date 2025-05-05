package main

import (
	"fmt"
	"math"
)

func checkPrime(num int) {
	if num == 2 {
		fmt.Println("num is prime")
		return
	}
	for i := 2; i < int(math.Sqrt(float64(num))+1); i++ {
		if num%i == 0 {
			fmt.Println("num is not prime")
			return
		}
	}
	fmt.Println("num is prime")

}

func main() {
	checkPrime(9)
}
