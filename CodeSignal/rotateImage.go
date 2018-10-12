package main

import "fmt"

func main() {
	// Test case 1
	a := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// Test case 2
	// b := [][]int{
	// 	{1, 2, 3, 4},
	// 	{5, 6, 7, 8},
	// 	{9, 10, 11, 12},
	// 	{13, 14, 15, 16},
	// }
	rotateImage(a)
}

// newImage[matrixSize - j - 1][matrixSize - 1] = a[0][0]

func rotateImage(a [][]int) [][]int {
	matrixSize := len(a)
	fmt.Println("flibidy", a[matrixSize-1][0])

	// counter := 0
	for i := 0; i < matrixSize/2; i++ {
		for j := 0; j <= matrixSize/2; j++ {
			// Handle the 4 corners
			// Counter should only be equivalent once per j
			if i == j {
				fmt.Println("Corners I:", i)
				fmt.Println("Corners J:", j)
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
				a[i][i] = temp2
			} else {
				// Rotate inner numbers
				fmt.Println("Middle I:", i)
				fmt.Println("Middle J:", j)
				// Store all 4 numbers
				top := a[i][j]
				fmt.Println("Sanity Check TOP", top)
				right := a[j][matrixSize-1]
				fmt.Println("Sanity Check RIGHT", right)
				bottom := a[matrixSize-1][j]
				fmt.Println("Sanity Check BOTTOM", bottom)
				left := a[j][i]
				fmt.Println("Sanity Check LEFT", left)
				// Replace all sides
				a[i][j] = left
				// Right
				a[j][matrixSize-1] = top
				// Bottom
				a[matrixSize-1][j] = right
				// Left
				a[j][i] = bottom

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
