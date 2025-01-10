package questions

/*
Rotate Matrix: Given an image represented by an NxN matrix, where each pixel in the image is 4
bytes, write a method to rotate the image by 90 degrees.
*/

func MatrixRotation(matrix [][]int) bool {
	// get length of the matrix
	n := len(matrix)

	// Check if the matrix is empty
	if n == 0 {
		return false
	}

	for layer := 0; layer < n/2; layer++ {
		first := layer
		last := n - 1 - layer
		for i := first; i < last; i++ {
			offset := i - first
			//save top
			top := matrix[first][i]

			// left to top
			matrix[first][i] = matrix[last-offset][first]

			// bottom to left
			matrix[last-offset][first] = matrix[last][last-offset]

			// right to bottom
			matrix[last][last-offset] = matrix[i][last]

			// top to right
			matrix[i][last] = top
		}
	}
	return true
}
