package slice

import (
	"fmt"
	"strings"
)

func Print1DSlice(slice []int) {
	cellWidth := 4
	for i := 0; i < len(slice); i++ {
		fmt.Printf("%*d", cellWidth, slice[i])
	}
}

func Print2DSlice(slice [][]int) {
	cellWidth := 4
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice[i]); j++ {
			if slice[i][j] == 0 {
				fmt.Print(strings.Repeat(" ", cellWidth))
			} else {
				fmt.Printf("%*d", cellWidth, slice[i][j])
			}
		}
		fmt.Println()
	}
}
