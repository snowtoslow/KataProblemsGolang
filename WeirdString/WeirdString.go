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
		if i%2==0 {
			log.Printf("IF:%d--->%s",i,string( unicode.ToUpper(myRune[i])))
			if myRune[i]==32 {
				log.Print("cont")
				continue
			}

			myRune[i] =  unicode.ToUpper(myRune[i])
		}else if i%2!=0{
			log.Printf("ELSE IF:%d--->%s",i,string(unicode.ToLower(myRune[i])))
			if myRune[i]==32 {
				continue
			}
			myRune[i] =  unicode.ToLower(myRune[i])
		}
	}
	return string(myRune)
}