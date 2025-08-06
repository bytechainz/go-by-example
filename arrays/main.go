package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp", a) // defalut [0 0 0 0 0]

	a[4] = 100
	fmt.Println("set :", a)
	fmt.Println("get:", a[4]) // 100

	fmt.Println("len : ", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl : ", b)

	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	/*
		Index	Nilai	Penjelasan
		0		100		nilai pertama, tidak pakai index
		1		0 		(default)	tidak ditentukan → default = 0
		2		0 		(default)	tidak ditentukan → default = 0
		3		400		index 3 ditentukan langsung
		4		500		setelah index 3

	*/

	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	b = [...]int{2: 500, 100, 90}
	fmt.Println("idx:", b)

	var matrix [2][3]int

	// Mengisi nilai
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[0][2] = 3
	matrix[1][0] = 4
	matrix[1][1] = 5
	matrix[1][2] = 6

	// Menampilkan array
	fmt.Println(matrix) // [[1 2 3] [4 5 6]]

	var twoDimensi [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoDimensi[i][j] = i + j
		}
	}
	fmt.Println("2d : ", twoDimensi)

	two2d := [2][3]int{
		{1, 2, 4},
		{4, 5, 6},
	}

	fmt.Println("2d : ", two2d)

}
