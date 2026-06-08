package main

import (
	"fmt"
)

func sort(s []int, x int) int {
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	halfLenght := len(s) / 2
	mid := s[halfLenght]
	if x != mid {
		if x < mid {
			for i := 0; i < (halfLenght); i++ {
				if x == s[i] {
					return i
				}
			}

		} else {
			for j := halfLenght + 1; j < len(s); j++ {
				if x == s[j] {
					return j
				}

			}
		}

	} else {
		return halfLenght
	}
	return -1
	
}

func main() {
	arr := []int{4, 2, 5, 1, 6, 7,32, 14, 12,14}
	s := sort(arr, 32)
	fmt.Println(s)
	// fmt.Println(r)
}
