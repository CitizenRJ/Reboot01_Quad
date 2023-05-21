package piscine

import "fmt"

func QuadC(x, y int) {
	if x < 1 || y < 1 {
		return
	}
	//top part
	fmt.Print("A")
	for i := 0; i < x-2; i++ {
		fmt.Print("B")
	}
	if x > 1 {
		fmt.Print("A")
	}
	fmt.Println("")
	//mid part
	if y > 2 {
		for j := 0; j < y-2; j++ {
			fmt.Print("B")
			for i := 0; i < x-2; i++ {
				fmt.Print(" ")
			}
			if x > 1 {
				fmt.Print("B")
			}
			fmt.Println("")
		}
	}
	// bottom part
	if y > 1 {
		fmt.Print("C")
		for i := 0; i < x-2; i++ {
			fmt.Print("B")
		}
		if x > 1 {
			fmt.Println("C")
		}
	}
}
