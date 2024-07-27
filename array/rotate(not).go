package array

import "fmt"

func Rotate(arr [][]int) [][]int {

	temp, length := 0, len(arr)

	for i := 0; i < length; i++ {
		for j := 0; j < len(arr[i]); j++ {
			temp = arr[i][length-j-1]
			arr[i][length-1] = arr[i][j]
		}
	}

	fmt.Printf("%d", temp)
	return arr
}
