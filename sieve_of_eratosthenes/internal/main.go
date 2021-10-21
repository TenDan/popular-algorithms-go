package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var max int

	for true {
		fmt.Print("Wprowadź maksymalną liczbę zakresu: ")
		maxString, _ := reader.ReadString('\n')

		RemoveEmptySpacesFromString(&maxString)

		var errMax error

		max, errMax = strconv.Atoi(maxString)

		if errMax == nil {
			break
		}
	}

	sieve := make([]bool, max+1)

	FillBooleanArrayWithTrueValues(&sieve)

	sieve[0] = false
	sieve[1] = false

	for i := 2; i < int(math.Sqrt(float64(max))); i++ {
		if sieve[i] {
			for j := int(math.Pow(float64(i), 2)); j <= max; j += i {
				sieve[j] = false
			}
		}
	}

	PrintPrimeNumbers(sieve)
}

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
