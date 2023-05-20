package piscine

import "fmt"

func QuadA(x,y int) {

	if y > 0 && x > 0 {

	//top part
		fmt.Print("o")
		for i := 0; i < x-2; i++ {
			fmt.Print("-")
		}
	if x > 1{
		fmt.Print("o")
	}
	fmt.Println("")	
	
	//mid part
	if y > 2 {
		for j := 0; j < y-2; j++ {
			fmt.Print("|")
			for i := 0; i < x-2; i++ {
				fmt.Print(" ")
				}
	if x > 1 {
		fmt.Print("|")	
	}
		fmt.Println("")
		}
	}
	
	// bottom part
	if y > 1{
	fmt.Print("o")
	for i := 0; i < x-2; i++ {
		fmt.Print("-")
	}
	if x > 1{
		fmt.Print("o")
	}
	fmt.Println("")
	}	
	}
}
