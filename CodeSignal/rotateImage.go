package main

import "fmt"

func main() {
	a := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotateImage(a)
}

// newImage[matrixSize - j - 1][matrixSize - 1] = a[0][0]

func rotateImage(a [][]int) [][]int {
	matrixSize := len(a)
	fmt.Println("flibidy", a[matrixSize-1][0])

	counter := 0
	for i := 0; i < matrixSize; i++ {
		for j := 0; j < matrixSize; j++ {
			// Handle the 4 corners
			// Counter should only be equivalent once per j
			if i == counter && j == counter {
				temp1 := 0
				temp2 := 0
				// Store the top left corner
				temp1 = a[0][0]
				// Store the top right corner
				temp2 = a[0][matrixSize-1]
				// Replace top right with top left
				a[0][matrixSize-1] = temp1
				// Store the bottom right corner
				temp1 = a[matrixSize-1][matrixSize-1]
				// Replace bottom right with top right
				a[matrixSize-1][matrixSize-1] = temp2
				// Store bottom left corner
				temp2 = a[matrixSize-1][0]
				// Replace bottom left
				a[matrixSize-1][0] = temp1
				// Replace top left with bottom left
				a[0][0] = temp2
				counter++
			} else if j <= j/2 {
				// fmt.Println("Sanity Check", a[i][j], a[j][i])
				// Middle terms only need to be moved 1/2 of length
				// Move the middle terms around
				// temp1 := a[i][j]
				// temp2 := a[j][matrixSize-1]
				// // fmt.Println("fuck", temp1)
				// // fmt.Println("shit", temp2)
				// // Replace subsequent numbers
				// // Right middle
				// a[j][matrixSize-1] = temp1
				// // Bottom middle
				// temp1 = a[matrixSize-1][j]
				// fmt.Println("fuck", temp1)

				// a[matrixSize-1][j] = temp2
				// fmt.Println("shit", temp2)

				// // Left middle
				// temp2 = a[i][j]

			}

		}
	}
	fmt.Println("Old Image:", a)

	return a
}

//
// 0, len / len, len / len, 0 / 0, 0
// a[i][j], a[i] = a[len][j]
// a = [[1,   2,   3, 4],
//      [5,   6,   7, 8],
//      [9,  10, 11, 12],
//      [13, 14, 15, 16]]
// if j == matrixSize-1 {
// 	temp := a[i][j]
// 	a[j][matrixSize-i-1] = temp
// } else {
// 	// a[j][matrixSize-1-j] = a[i][j]

// }
