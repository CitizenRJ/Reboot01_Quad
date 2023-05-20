package piscine

import "fmt"

func QuadB(x, y int) {
	if x < 0 || y < 0 {
		return
	}
	// rows
	for i := 0; i < y; i++ {
		// columons
		for j := 0; j < x; j++ {
			// top
			if i == 0 || i == y-1 {
				if j == 0 && i == 0 {
					fmt.Print("/")
				} else if j == x-1 && i == 0 {
					fmt.Print("\\")
				} else if j == 0 && i == y-1 {
					fmt.Print("\\")
				} else if j == x-1 && i == y-1 {
					fmt.Print("/")
				} else {
					// top
					fmt.Print("*")
				}
			} else {
				// space inside
				if j == 0 || j == x-1 {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Print("\n")
	}
}
