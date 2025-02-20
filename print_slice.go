package print_slice

import "fmt"

func PrintArray2D(array [][]int) {
	for _, rows := range array {
		for _, cols := range rows {
			fmt.Printf("%d ", cols)
		}
		fmt.Println()
	}
}
