package main

import (
	"fmt"
)

func main() {
	const usdToEur = 0.85
	const usdToRub = 78.5

	eurTORub := fmt.Sprintf("%.2f", 1/usdToEur*usdToRub)
	fmt.Println("Euro to Ruble: ", eurTORub)
}
