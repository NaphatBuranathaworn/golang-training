package calculator

// Calculate something
func Calculate(arr [][]int) int {
	result := 0
	x := 0
	y := 0

	z := len(arr) - 1

	for i := 0; i < len(arr); i++ {
		x = x + arr[i][i]
		y = y + arr[z][i]
		z--
	}

	result = x - y;
	if (result < 0) {
        result = result * -1
    }

	return result;
}