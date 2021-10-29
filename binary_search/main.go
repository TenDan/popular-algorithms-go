package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	scripts "github.com/TenDan/popular-algorithms-go/binary_search/internal"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Wprowadź ciąg liczb (liczby mają być oddzielone przecinkami \",\")")

	line, _ := reader.ReadString('\n')

	scripts.RemoveWhiteSpaces(&line)

	numberStrings := strings.Split(line, ",")

	numbers := make(scripts.IntArray, len(numberStrings))

	for i, v := range numberStrings {
		var err error
		for true {
			numbers[i], err = strconv.Atoi(v)
			if err == nil {
				break
			}
		}
	}

	sort.Ints(numbers)

	numbers.Print(", ")

	fmt.Print("Szukana liczba: ")

	search, _ := reader.ReadString('\n')

	scripts.RemoveWhiteSpaces(&search)

	var (
		searchNumber int
		err          error
	)

	for true {
		searchNumber, err = strconv.Atoi(search)
		if err == nil {
			break
		}
	}

	fmt.Printf("Element na indeksie %d", scripts.SearchForNumber(numbers, 0, len(numbers), searchNumber))
}
