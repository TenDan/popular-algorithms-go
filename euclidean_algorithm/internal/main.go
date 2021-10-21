package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var (
		firstNumber  = 0
		secondNumber = 0
	)

	for true {
		fmt.Print("Wprowadź pierwszą liczbę: ")
		first, _ := reader.ReadString('\n')

		fmt.Print("Wprowadź drugą liczbę: ")
		second, _ := reader.ReadString('\n')

		//fmt.Println(first, second)

		RemoveEmptySpacesFromString(&first)
		RemoveEmptySpacesFromString(&second)

		//fmt.Println(first, second)

		var (
			errFirst  error
			errSecond error
		)

		firstNumber, errFirst = strconv.Atoi(first)
		secondNumber, errSecond = strconv.Atoi(second)

		//fmt.Println(firstNumber, secondNumber)
		//fmt.Println(errFirst.Error(), errSecond.Error())

		if errFirst == nil && errSecond == nil {
			break
		}
	}

	gcdByComparison := GcdByComparison(firstNumber, secondNumber)
	gcdByModulo := GcdByModulo(firstNumber, secondNumber)
	gcdRecursive := GcdRecursive(firstNumber, secondNumber)
	gcdRecursiveComparison := GcdRecursiveComparison(firstNumber, secondNumber)

	fmt.Printf("Największy wspólny dzielnik dla liczb: %d i %d to %d\n", firstNumber, secondNumber, gcdByComparison)
	fmt.Printf("Największy wspólny dzielnik dla liczb: %d i %d to %d\n", firstNumber, secondNumber, gcdByModulo)
	fmt.Printf("Największy wspólny dzielnik dla liczb: %d i %d to %d\n", firstNumber, secondNumber, gcdRecursive)
	fmt.Printf("Największy wspólny dzielnik dla liczb: %d i %d to %d\n", firstNumber, secondNumber, gcdRecursiveComparison)
}

func GcdByComparison(first int, second int) int {
	for first != second {
		if first > second {
			first = first - second
		} else {
			second = second - first
		}
	}
	return first
}

func GcdByModulo(first int, second int) int {
	for second != 0 {
		second, first = first%second, second
	}
	return first
}

func GcdRecursive(first int, second int) int {
	if second == 0 {
		return first
	} else {
		return GcdRecursive(second, first%second)
	}
}

func GcdRecursiveComparison(first int, second int) int {
	if first == second {
		return first
	} else {
		if first > second {
			return GcdRecursiveComparison(first-second, second)
		} else {
			return GcdRecursiveComparison(first, second-first)
		}
	}
}

func RemoveEmptySpacesFromString(str *string) {
	*str = strings.ReplaceAll(*str, "\n", "")
	*str = strings.ReplaceAll(*str, "\r", "")
	*str = strings.ReplaceAll(*str, " ", "")
}
