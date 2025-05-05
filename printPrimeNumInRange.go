package main

import (
	"fmt"
	"math"
)

func printPrime(num int) {
	var num1 []int
	for j := 2; j <= num; j++ {
		var flag = true
		if num == 2 {
			flag = true
		}
		for i := 2; i < int(math.Sqrt(float64(j))+1); i++ {
			if j%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			num1 = append(num1, j)
		}
	}
	fmt.Println(num1)
}
func main() {
	printPrime(9)
}
