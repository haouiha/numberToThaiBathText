package main

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestNumberToThaiBathText(t *testing.T) {
	testCases := []struct {
		input    decimal.Decimal
		expected string
	}{
		{decimal.NewFromFloat(0.00), "ศูนย์บาทถ้วน"},
		{decimal.NewFromFloat(0), "ศูนย์บาทถ้วน"},
		{decimal.NewFromFloat(0.25), "ยี่สิบห้าสตางค์"},
		{decimal.NewFromFloat(0.500), "ห้าสิบสตางค์"},
		{decimal.NewFromFloat(0.6), "หกสิบสตางค์"},
		{decimal.NewFromFloat(0.75), "เจ็ดสิบห้าสตางค์"},
		{decimal.NewFromFloat(1), "หนึ่งบาทถ้วน"},
		{decimal.NewFromFloat(10), "สิบบาทถ้วน"},
		{decimal.NewFromFloat(20.00), "ยี่สิบบาทถ้วน"},
		{decimal.NewFromFloat(21), "ยี่สิบเอ็ดบาทถ้วน"},
		{decimal.NewFromFloat(100), "หนึ่งร้อยบาทถ้วน"},
		{decimal.NewFromFloat(101), "หนึ่งร้อยเอ็ดบาทถ้วน"},
		{decimal.NewFromFloat(1000), "หนึ่งพันบาทถ้วน"},
		{decimal.NewFromFloat(1234), "หนึ่งพันสองร้อยสามสิบสี่บาทถ้วน"},
		{decimal.NewFromFloat(10000), "หนึ่งหมื่นบาทถ้วน"},
		{decimal.NewFromFloat(33333.75), "สามหมื่นสามพันสามร้อยสามสิบสามบาทเจ็ดสิบห้าสตางค์"},
		{decimal.NewFromFloat(100000), "หนึ่งแสนบาทถ้วน"},
		{decimal.NewFromFloat(1000000), "หนึ่งล้านบาทถ้วน"},
		{decimal.NewFromFloat(1234567), "หนึ่งล้านสองแสนสามหมื่นสี่พันห้าร้อยหกสิบเจ็ดบาทถ้วน"},
		{decimal.NewFromFloat(1000000000000), "หนึ่งล้านล้านบาทถ้วน"},
	}

	for _, tc := range testCases {
		actual := numberToThaiBathText(tc.input)
		if actual != tc.expected {
			t.Errorf("ConvertBahtText(%s) = %s; expected %s", tc.input.String(), actual, tc.expected)
		}
	}
}
