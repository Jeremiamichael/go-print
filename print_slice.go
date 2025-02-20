package print_slice

import "fmt"

func PrintSlice(slice []int) {
	for i := 0; i < len(slice); i++ {
		fmt.Print(slice[i], " ")
	}
}

func PrintSlice2D(slice [][]int) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice[i]); j++ {
			fmt.Print(slice[i][j], "\t")
		}
		fmt.Println()
	}
}
