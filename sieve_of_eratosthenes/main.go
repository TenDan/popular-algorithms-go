package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"

	scripts "github.com/TenDan/popular-algorithms-go/sieve_of_eratosthenes/internal"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var max int

	for true {
		fmt.Print("Wprowadź maksymalną liczbę zakresu: ")
		maxString, _ := reader.ReadString('\n')

		scripts.RemoveEmptySpacesFromString(&maxString)

		var errMax error

		max, errMax = strconv.Atoi(maxString)

		if errMax == nil {
			break
		}
	}

	sieve := make([]bool, max+1)

	scripts.FillBooleanArrayWithTrueValues(&sieve)

	sieve[0] = false
	sieve[1] = false

	for i := 2; i < int(math.Sqrt(float64(max))); i++ {
		if sieve[i] {
			for j := int(math.Pow(float64(i), 2)); j <= max; j += i {
				sieve[j] = false
			}
		}
	}

	scripts.PrintPrimeNumbers(sieve)
}
