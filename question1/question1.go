package que1

import (
	"strconv"
	"strings"
)

func contain(commas *[]int, index int) bool {
	for _, value := range *commas {
		if value == index {
			return true
		}
	}
	return false
}

func findDigitIndex(s []string) int {
	for i := 0; i < len(s); i++ {
		if s[i] == "." {
			return i - 1
		}
	}
	return len(s) - 1
}

func commaPosition(c *[]int, index int) {
	for index >= 3 {
		index -= 3
		*c = append(*c, index)
	}
}

func f(num float64) string {
	b := strings.Split(strconv.FormatFloat(num, 'f', -1, 64), "")
	index := findDigitIndex(b)

	commas := []int{}
	commaPosition(&commas, index)

	ans := []string{}
	for j := 0; j < len(b); j++ {
		ans = append(ans, b[j])
		if contain(&commas, j) && b[j] != "-" {
			ans = append(ans, ",")
		}
	}
	return strings.Join(ans, "")
}

// func main() {
// 	fmt.Printf("F(9527) => %s\n", f(9527))
// 	fmt.Printf("F(3345678) => %s\n", f(3345678))
// 	fmt.Printf("F(-1234.45) => %s\n", f(-1234.45))

// }
