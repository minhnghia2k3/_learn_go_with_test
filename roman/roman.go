package main

import (
	"strings"
)

type RomanNumeral struct {
	Number int
	Symbol string
}

var sliceOfRomanNumeral = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for _, s := range sliceOfRomanNumeral {
		for arabic >= s.Number {
			result.WriteString(s.Symbol)
			arabic -= s.Number
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) int {
	total := 0

	// IV
	// res = 1 => V
	// res = 6 => ""

	// VII
	// res = 5 => II
	// res = 6 => I
	// res = 7 = > I

	for _, s := range sliceOfRomanNumeral {
		for strings.HasPrefix(roman, s.Symbol) {
			total += s.Number
			roman = strings.TrimPrefix(roman, s.Symbol)
		}
	}
	return total
}
