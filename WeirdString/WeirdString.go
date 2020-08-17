package main

import (
	"fmt"
	"log"
	"unicode"
)

func main() {


	myString := "This is a test Looks like you passed"

	fmt.Println(weirdCase(myString))
}

func weirdCase(inputString string) string{
	myRune:= []rune(inputString)
	for i:=0;i<len(myRune);i++{
		if myRune[i]==32 {
			continue
		}else if i%2==0{
			myRune[i] =  unicode.ToUpper(myRune[i])
		} else {
			myRune[i] =  unicode.ToLower(myRune[i])
		}
	}
	return string(myRune)
}