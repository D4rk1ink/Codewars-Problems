package main

import (
	"math"
	"strconv"
)

func Add(num1 string, num2 string) string {
	lenNum1 := len(num1)
	lenNum2 := len(num2)
	result := ""
	pivot := 1
	extended := 0

	for {
		if pivot > lenNum1 && pivot > lenNum2 {
			break
		}

		n1 := 0
		n2 := 0
		if pivot <= lenNum1 {
			n1, _ = strconv.Atoi(string(num1[lenNum1-pivot]))
		}
		if pivot <= lenNum2 {
			n2, _ = strconv.Atoi(string(num2[lenNum2-pivot]))
		}
		cal := n1 + n2 + extended
		result = strconv.Itoa(cal%10) + result
		extended = int(math.Floor(float64(cal / 10)))

		pivot += 1
	}

	if extended != 0 {
		result = strconv.Itoa(extended) + result
	}

	return result
}
