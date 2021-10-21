package scripts

import (
	"fmt"
	"strings"
)

func FillBooleanArrayWithTrueValues(arr *[]bool) {
	for i := 0; i < len(*arr); i++ {
		(*arr)[i] = true
	}
}

func PrintPrimeNumbers(arr []bool) {
	firstNumber := true
	for i := 0; i < len(arr); i++ {
		if arr[i] {
			if firstNumber {
				firstNumber = false
			} else {
				fmt.Print(", ")
			}
			fmt.Print(i)
		}
	}
}

func RemoveEmptySpacesFromString(str *string) {
	*str = strings.ReplaceAll(*str, "\r", "")
	*str = strings.ReplaceAll(*str, "\n", "")
	*str = strings.ReplaceAll(*str, " ", "")
}
