package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	myString := "A"
	arrayOfStrings:=splitStringByWhiteSpaces(myString)
	fmt.Println(cipher(arrayOfStrings,create2DArrayOfRunes,swapFirstLastAndGetFirstLetters))
}



func splitStringByWhiteSpaces(input string) []string {
	var sliceOfStrings []string
	sliceOfStrings = strings.Fields(input)

	return sliceOfStrings
}

func create2DArrayOfRunes(inputArray []string) [][]rune {
	var myArray [][]rune
	for _,k := range inputArray{
		myArray = append(myArray, []rune(k))
	}
	return myArray
}


func swapFirstLastAndGetFirstLetters(inputArray [][]rune) [][]rune {

	for _,v := range inputArray{
		if len(v)==1 {
			continue
		}
		v[1],v[len(v)-1] = v[len(v)-1], v[1]
	}

	return inputArray
}

func cipher(input []string, getRunesArray func(input []string) [][]rune, doSwapAndGetFirst func(myArray [][]rune) [][]rune) string {
	swappedArray := doSwapAndGetFirst(getRunesArray(input))
	var resultArray string

	for k,v :=range swappedArray{
		if len(v)==1 && k == len(swappedArray)-1 {
			resultArray += formatInt32(v[0])
		}else if len(v)==1  {
			resultArray += formatInt32(v[0]) + " "
		} else if k == len(swappedArray)-1 {
			resultArray += formatInt32(v[0])+string(v[1:])
		} else{
			resultArray += formatInt32(v[0])+string(v[1:]) + " "
		}

	}

	return resultArray

}

func formatInt32(n int32) string {
	return strconv.FormatInt(int64(n), 10)
}








