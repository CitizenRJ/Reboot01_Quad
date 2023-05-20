package piscine

import "fmt"

func QuadD(x, y int) {
	if x < 0 || y < 0 {
		return
	}
	// rows
	for i := 0; i < y; i++ {
		// cols
		for j := 0; j < x; j++ {
			// corners
			if i == 0 || i == y-1 {
				if j == 0 && i == 0 {
					fmt.Print("A")
				} else if j == x-1 && i == 0 {
					fmt.Print("C")
				} else if j == 0 && i == y-1 {
					fmt.Print("A")
				} else if j == x-1 && i == y-1 {
					fmt.Print("C")
				} else {
					// inner
					fmt.Print("B")
				}
			} else {
				// empty inside
				if j == 0 || j == x-1 {
					fmt.Print("B")
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Print("\n")
	}
}
