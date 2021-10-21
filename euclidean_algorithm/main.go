package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	scripts "github.com/TenDan/popular-algorithms-go/euclidean_algorithm/internal"
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

		scripts.RemoveEmptySpacesFromString(&first)
		scripts.RemoveEmptySpacesFromString(&second)

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

	gcdByComparison := scripts.GcdByComparison(firstNumber, secondNumber)
	gcdByModulo := scripts.GcdByModulo(firstNumber, secondNumber)
	gcdRecursive := scripts.GcdRecursive(firstNumber, secondNumber)
	gcdRecursiveComparison := scripts.GcdRecursiveComparison(firstNumber, secondNumber)

	fmt.Printf("Największy wspólny dzielnik dla liczb: %d i %d to %d\n", firstNumber, secondNumber, gcdByComparison)
	fmt.Printf("Największy wspólny dzielnik dla liczb: %d i %d to %d\n", firstNumber, secondNumber, gcdByModulo)
	fmt.Printf("Największy wspólny dzielnik dla liczb: %d i %d to %d\n", firstNumber, secondNumber, gcdRecursive)
	fmt.Printf("Największy wspólny dzielnik dla liczb: %d i %d to %d\n", firstNumber, secondNumber, gcdRecursiveComparison)
}
