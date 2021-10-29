package internal

import (
	"fmt"
	"strings"
)

func RemoveWhiteSpaces(s *string) {
	*s = strings.ReplaceAll(*s, " ", "")
	*s = strings.ReplaceAll(*s, "\n", "")
	*s = strings.ReplaceAll(*s, "\r", "")
	*s = strings.ReplaceAll(*s, "\t", "")
	*s = strings.ReplaceAll(*s, "\b", "")
}

type IntArray []int

func (a IntArray) Print(sep string) {
	fmt.Print("[")
	i := 0
	for _, v := range a {
		if i > 0 {
			fmt.Print(sep)
		}
		fmt.Printf("%d", v)
		i++
	}
	fmt.Print("]\n")
}

func SearchForNumber(arr []int, low, high, target int) int {
	if low > high {
		return -1
	}

	mid := (low + high) / 2

	if target == arr[mid] {
		return mid
	} else if target < arr[mid] {
		return SearchForNumber(arr, low, mid-1, target)
	} else {
		return SearchForNumber(arr, mid+1, high, target)
	}
}
