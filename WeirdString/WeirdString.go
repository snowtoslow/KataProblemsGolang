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
			log.Print("CONTINUE")
			continue
		}else if i%2==0{
			log.Printf("EVEN:%d--->%s",i,string(unicode.ToUpper(myRune[i])))
			myRune[i] =  unicode.ToUpper(myRune[i])
		} else {
			log.Printf("ODD:%d--->%s",i, string(unicode.ToLower(myRune[i])))
			myRune[i] =  unicode.ToLower(myRune[i])
		}
	}
	return string(myRune)
}