package slice

import "fmt"

func print1DSlice(slice []int) {
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
}

func print2DSlice(slice [][]int) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice[i]); j++ {
			fmt.Println(slice[i][j])
		}
	}
}
