package main

import "fmt"

func counter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

func multiplier(factor int) func(int) int {
	return func(x int) int {
		return x + factor
	}
}

func main() {
	c := counter()

	fmt.Println(c()) // 1
	fmt.Println(c()) // 2
	fmt.Println(c()) // 3

	d := counter()   // closure baru, count dimulai lagi dari 0
	fmt.Println(d()) // 1

	double := multiplier(2)
	triple := multiplier(3)

	fmt.Println(double(5))
	fmt.Println(triple(5))

}
