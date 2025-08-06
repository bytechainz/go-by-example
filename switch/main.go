package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("write ", i, " as ")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	fmt.Println(time.Now().Weekday())

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: // dengan 2 kondisi
		fmt.Println("its weekend")
	default:
		fmt.Println("its a weekday")
	}

	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("its before noon")
	default:
		fmt.Println("its after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("i m a bool")
		case int:
			fmt.Println("im an int")
		default:
			fmt.Println("dont know type", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
