package main

func main() {

	multiplicationTable(1)

}

func multiplicationTable(size int) [][]int {
	myArray := make([][]int, size)
	for i:= range myArray{
		myArray[i] = make([]int, size)
		vector := make([]int, size)
		for j := range myArray[i]{
			vector[j] = (i+1)*(j+1)
			myArray[i][j] = vector[j]
		}
	}
	return myArray

}