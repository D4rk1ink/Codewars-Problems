package main

import (
	"fmt"
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

func Multiply(multiplicand string, multiplier string) string {
	if multiplicand == "0" || multiplier == "0" {
		return "0"
	}
	sum := ""
	lenMultiplicand := len(multiplicand)
	lenMultiplier := len(multiplier)

	for i := lenMultiplier - 1; i >= 0; i-- {
		m1, _ := strconv.Atoi(string(multiplier[i]))
		if m1 == 0 {
			continue
		}
		extended := 0
		lastNonZeroIndex := lenMultiplier - 1
		result := ""
		for z := 0; z < (lenMultiplier - 1 - i); z++ {
			result += "0"
		}
		for j := lenMultiplicand - 1; j >= 0; j-- {
			m2, _ := strconv.Atoi(string(multiplicand[j]))
			cal := m1*m2 + extended

			if m2 == 0 && extended != 0 {
				result = strconv.Itoa(cal) + result
				extended = 0
				lastNonZeroIndex = j
			} else {
				result = strconv.Itoa(cal%10) + result
				extended = int(math.Floor(float64(cal / 10)))
			}
			if m2 != 0 {
				lastNonZeroIndex = j
			}
		}
		if extended != 0 {
			result = strconv.Itoa(extended) + result
		}

		sum = Add(result[lastNonZeroIndex:], sum)
	}

	return sum
}

func TestMultiply() {
	testcases := [][]string{
		{"2", "0", "0"},
		{"0", "30", "0"},
		{"0000001", "3", "3"},
		{"1000001", "3", "3000003"},
		{"2", "3", "6"},
		{"29", "5", "145"},
		{"30", "69", "2070"},
		{"11", "85", "935"},
		{"1009", "03", "3027"},
		{
			"01857676599633787571309758335143087859263372693536636596716263717576954",
			"9",
			"16719089396704088141787825016287790733370354241829729370446373458192586",
		},
		{"018", "9", "162"},
		{
			"1020303004875647366210",
			"2774537626200857473632627613",
			"2830869077153280552556547081187254342445169156730",
		},
		{
			"58608473622772837728372827",
			"7586374672263726736374",
			"444625839871840560024489175424316205566214109298",
		},
		{
			"01857676599633787571309758335143087859263372693536636596716263717576954",
			"60197671035752796144013026738956372679497185647413406253210790599521935",
			"111827804835570597469803982908218785104176557578866304288908024762591625515223866705731753296182036381812491317890763876470478065688973485990",
		},
	}

	for _, testcase := range testcases {
		result := Multiply(testcase[0], testcase[1])
		if result == testcase[2] {
			fmt.Println("Passes")
		} else {
			fmt.Printf("Failed, expect %s but got %s\n", testcase[2], result)
		}
	}
}
