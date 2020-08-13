package main

import (
	"fmt"
	"unicode"
)

func main() {

	myString := "Weird string"

	fmt.Println(weirdCase(myString))

	fmt.Println(test(weirdCase(myString),"WeIrD StRiNg"))


}

func weirdCase(inputString string) string{
	myRune:= []rune(inputString)
	for i:=0;i<len(myRune);i++{
		if i%2==0 {
			if myRune[i]==32 {
				continue
			}
			myRune[i] =  unicode.ToUpper(myRune[i])
		}
	}
	return string(myRune)
}

func test(s string, r string) bool{
	return s == r
}