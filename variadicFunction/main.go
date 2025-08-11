package main

import "fmt"

func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 1)
	sum(1, 2, 3, 4)

	nums := []int{1, 2, 3, 4, 5}
	sum(nums...)
}

/*


Variadic functions can be called with
any number of trailing arguments. For example,
 fmt.Println is a common variadic function.
*/
