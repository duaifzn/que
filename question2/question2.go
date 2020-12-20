package que2

// import (
// 	"fmt"
// )

func pipe(num int, callbacks ...func(*int, int)) int {
	for _, callback := range callbacks {
		callback(&num, 1)
	}
	return num
}

func increment(a *int, b int) {
	*a += b
}

// func main() {
// 	fmt.Printf("pipe(5, increment) = %d\n", pipe(5, increment))
// 	fmt.Printf("pipe(5, increment, increment, increment) = %d\n", pipe(5, increment, increment, increment))
// }
