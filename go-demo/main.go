package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2

	userHeight := 1.8
	userKg := 100.0

	imt := userKg / math.Pow(userHeight, IMTPower)

	fmt.Print(imt)
}
