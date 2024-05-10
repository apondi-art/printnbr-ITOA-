package main

import (
	"fmt"
)

//function takes integers and return strings
//it uses division which is div to update value of t
//the mod of division is converted back to string(getting asccii representation of each mod =(int32(mod)+'0')= int 32 is rune presentation
//then after changing to rune ,append that rune value to []string(str)
//reverse loopover the slice then add to finalstring

func main() {
	s := 350
	fmt.Println(Itoa(s))
}
func Itoa(t int) string {
	var div int
	var mod int
	var run rune
	var str []string
	var finalstring string
	div = t
	for div > 0 {
		mod = div % 10
		div = div / 10

		run = int32(mod) + '0'
		str = append(str, string(run))

	}
	for i := len(str) - 1; i >= 0; i-- {
		finalstring += str[i]

	}
	return finalstring
}
