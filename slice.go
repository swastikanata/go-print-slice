package slice

import "fmt"

func Print1DSlice(slice []int) {
	for i := 0; i < len(slice); i++ {
		fmt.Printf("%d\t", slice[i])
	}
	fmt.Println()
}

func Print2DSlice(slice [][]int) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice[i]); j++ {
			if slice[i][j] == 0 {
				fmt.Print("\t")
			} else {
				fmt.Printf("%d\t", slice[i][j])
			}
		}
		fmt.Println()
	}
}
