package main

import (
	"fmt"
)

func main() {
	point1 := lengthOfLongestSubstring("eeydgwdykpv")
	fmt.Println(len("eeydgwdykpv"))
	fmt.Printf("%d", point1)
}

/*
	4ms
 */
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	maxlength := 0
	arr:=make([]rune,len(s))//alocate memory for all
	substringslice := arr[0:0]
	fmt.Printf("len:%d\r\n",len(substringslice))
	fmt.Printf("cap:%d\r\n",cap(substringslice))

	for _, char := range s {
		indexof := IndexOf(substringslice, char)
		if indexof > -1 {
			substringslice = substringslice[indexof+1:]
		}
		substringslice = append(substringslice, char)

		if maxlength < len(substringslice) {
			maxlength = len(substringslice)
		}
	}
	return maxlength
}

func IndexOf(runes []rune, target rune) int {
	if runes == nil {
		return -1
	}
	for index, value := range runes {
		if value == target {
			return index
		}
	}
	return -1
}

