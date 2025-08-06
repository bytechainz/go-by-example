package main

import (
	"fmt"
	"slices"
)

func main() {

	/*
		slice = make([]string, 3)
		slice bertipe string dengan panjang 3 dan kapasitas 3,

		Panjang		Jumlah elemen yang saat ini tersedia
		Kapasitas	Jumlah elemen maksimum yang bisa ditampung sebelum realokasi

		s := make([]int, 3, 5)
		fmt.Println(len(s)) // 3
		fmt.Println(cap(s)) // 5
	*/
	newSlice := make([]int, 3, 5)
	fmt.Println("len : ", len(newSlice))
	fmt.Println("cap : ", cap(newSlice))

	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0) // s == nil -> true

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set : ", s)
	fmt.Println("get:", s[2])

	fmt.Println("len : ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)
	fmt.Println("len2 : ", len(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy : ", c) // [a b c d e f]

	l := s[2:5]              // ambil nilai setelah element 2 sampai element ke 5
	fmt.Println("sli : ", l) // [c d e]

	l = s[:5]
	fmt.Println("sl2:", l) // [a b c d e]

	l = s[2:]
	fmt.Println("sl3:", l) // [c d e f]

	t := []string{"g", "h", "i"}
	fmt.Println("dcl : ", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	} else {
		fmt.Println("t != t2")
	}

	twoD := make([][]int, 3)

	fmt.Println(twoD[2])

	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d :", twoD)

}
