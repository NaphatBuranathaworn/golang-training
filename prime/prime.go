package prime

import (
	"fmt"
)

// PrintNumber prints prime number from 1 to max
func PrintNumber(max int) {
	for i := 1; i <= max; i++ {
		var count int
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				count++
			}
		}
		if count == 2 {
			fmt.Print(i, ",")
		}
	}
}