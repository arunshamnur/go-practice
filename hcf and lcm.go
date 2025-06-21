package main

import "fmt"

func findhcf(n1 int, n2 int) {
	var lowest int
	var hcf int
	if n1 > n2 {
		lowest = n2
	} else {
		lowest = n1
	}
	for i := 1; i < lowest; i++ {
		if n1%i == 0 && n2%i == 0 {
			hcf = i
		}
	}
	fmt.Println(hcf)

}

func findLcm(n1 int, n2 int) {
	var highest int
	var lcm int
	if n1 < n2 {
		highest = n1
	} else {
		highest = n2
	}

	for {
		if highest%n1 == 0 && highest%n2 == 0 {
			lcm = highest
			break
		}
		highest++
	}
	fmt.Println(lcm)

}

func main() {
	findhcf(8, 20)
	findLcm(8, 20)

}
