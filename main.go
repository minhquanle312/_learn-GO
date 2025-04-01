package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	investmentAmount, years, expectedResultRate := 1000.0, 10.0, 5.5

	futureValue := investmentAmount * math.Pow(1+expectedResultRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
