package scripts

import (
	"strings"
)

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
