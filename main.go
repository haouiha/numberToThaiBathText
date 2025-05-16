package main

import (
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
)

var (
	millionPos  = 6
	thaiNumbers = []string{"", "หนึ่ง", "สอง", "สาม", "สี่", "ห้า", "หก", "เจ็ด", "แปด", "เก้า"}
	thaiUnits   = []string{"", "สิบ", "ร้อย", "พัน", "หมื่น", "แสน"}
)

func main() {
	inputs := []decimal.Decimal{
		decimal.NewFromFloat(1234),
		decimal.NewFromFloat(33333.75),
	}

	for _, input := range inputs {
		fmt.Println("Input:", input)
		fmt.Println("Output:", numberToThaiBathText(input))
		fmt.Println("--------------")
	}
}

func numberToThaiBathText(amount decimal.Decimal) string {
	if amount.IsZero() {
		return "ศูนย์บาทถ้วน"
	}

	amountString := amount.String()
	stringParts := strings.Split(amountString, ".")

	integerPart := stringParts[0]

	text := convertNumberToThaiText(integerPart)

	if amount.IsInteger() {
		text += "บาทถ้วน"
	} else {
		fractionalPart := stringParts[1]
		if len(fractionalPart) == 1 {
			fractionalPart += "0"
		} else if len(fractionalPart) > 2 {
			fractionalPart = fractionalPart[:2]
		}

		if integerPart != "0" {
			text += "บาท"
		}

		text += convertNumberToThaiText(fractionalPart) + "สตางค์"
	}

	return text
}

func convertNumberToThaiText(num string) string {
	groups := splitIntoGroups(num)
	groupCount := len(groups)
	result := ""

	for i := 0; i < groupCount; i++ {
		text := groupStringToText(groups[i])
		result += text
		if i < groupCount-1 {
			result += "ล้าน"
		}
	}

	return result
}

func splitIntoGroups(num string) []string {
	var groups []string
	last := len(num)

	for last != 0 {
		start := last - millionPos
		if start < 0 {
			start = 0
		}
		groups = append([]string{num[start:last]}, groups...)

		last = start
	}

	return groups
}

func groupStringToText(group string) string {
	length := len(group)
	result := ""

	for i := 0; i < length; i++ {
		digit := group[i] - '0'
		pos := length - i - 1

		if pos == 0 && digit == 1 && length > 1 {
			result += "เอ็ด"
		} else if pos == 1 && digit == 2 {
			result += "ยี่"
		} else if pos == 1 && digit == 1 {
			//skip "หนึ่ง" หลัก สิบ
		} else {
			result += thaiNumbers[digit]
		}

		if digit != 0 {
			result += thaiUnits[pos]
		}

	}

	return result
}
