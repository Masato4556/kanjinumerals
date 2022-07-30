package kanjinumerals

import (
	"math"
)

var LargeNumeralSymbols = map[string]int{"万": 4, "億": 8, "兆": 12}
var MediumNumeralSymbols = map[string]int{"十": 1, "百": 2, "千": 3}
var SmallNumeralSymbols = map[string]int{"〇": 0, "一": 1, "二": 2, "三": 3, "四": 4, "五": 5, "六": 6, "七": 7, "八": 8, "九": 9}

type FourDigitKanji struct {
	V []string // 4桁の数値
	E string   // 10の冪乗
}
type FourDigitKanjis []FourDigitKanji
type FourDigitNumber struct {
	V int // 4桁の数値
	E int // 10の冪乗
}
type FourDigitNumbers []FourDigitNumber

func (k FourDigitKanji) ToFourDigitNumber() (n FourDigitNumber) {
	n = FourDigitNumber{V: k.numberV(), E: 0}
	if e, ok := LargeNumeralSymbols[k.E]; ok {
		n.E = e
	}
	return
}

func (k FourDigitKanji) numberV() (nv int) {
	temp := 0
	digits := []int{}
	for i, v := range k.V {
		if mns, ok := MediumNumeralSymbols[v]; ok {
			if i == 0 {
				temp = 1
			}
			temp *= int(math.Pow10(mns))
			digits = append(digits, temp)
			temp = 0
		}
		if sns, ok := SmallNumeralSymbols[v]; ok {
			temp += sns
		}
	}
	if temp != 0 {
		digits = append(digits, temp)
	}
	for _, v := range digits {
		nv += v
	}
	return nv
}

func (ks FourDigitKanjis) ToFourDigitNumbers() (ns FourDigitNumbers) {
	for _, k := range ks {
		ns = append(ns, k.ToFourDigitNumber())
	}
	return
}

func (ns FourDigitNumbers) ToInt() int {
	number := 0
	for _, n := range ns {
		number += n.V * int(math.Pow10(n.E))
	}
	return number
}
