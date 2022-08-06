package kanjinumerals

import (
	"math"
	"strings"
)

var LargePowerNumeralSymbols = map[string]int{"万": 4, "億": 8, "兆": 12}
var SmallPowerNumeralSymbols = map[string]int{"十": 1, "百": 2, "千": 3}
var ArabicNumeralSymbols = map[string]int{"〇": 0, "一": 1, "二": 2, "三": 3, "四": 4, "五": 5, "六": 6, "七": 7, "八": 8, "九": 9}

func findSmallPowerNumeralKanji(n int) string {
	for k, v := range SmallPowerNumeralSymbols {
		if v == n {
			return k
		}
	}
	return ""
}

func findArabicNumeralKanji(n int) string {
	for k, v := range ArabicNumeralSymbols {
		if v == n {
			return k
		}
	}
	return ""
}

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
	if e, ok := LargePowerNumeralSymbols[k.E]; ok {
		n.E = e
	}
	return
}

func (k FourDigitKanji) IncludeSmallPowerNumeralSymbols() bool {
	for symbol, _ := range SmallPowerNumeralSymbols {
		if contains(k.V, symbol) {
			return true
		}
	}
	return false
}

func (k FourDigitKanji) numberV() (nv int) {
	if k.IncludeSmallPowerNumeralSymbols() {
		return k.vToNumberWithPowers()
	}

	return k.vToNumberWithoutPowers()
}

func (k FourDigitKanji) vToNumberWithPowers() (nv int) {
	temp := 0
	digits := []int{}
	for _, v := range k.V {
		if mns, ok := SmallPowerNumeralSymbols[v]; ok {
			if temp == 0 {
				temp = 1
			}
			temp *= int(math.Pow10(mns))
			digits = append(digits, temp)
			temp = 0
		}
		if sns, ok := ArabicNumeralSymbols[v]; ok {
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

func (k FourDigitKanji) vToNumberWithoutPowers() (nv int) {
	for i, v := range k.V {
		if i > 0 {
			nv *= 10
		}
		if sns, ok := ArabicNumeralSymbols[v]; ok {
			nv += sns
		}
	}
	return nv
}

func (ks FourDigitKanjis) ToFourDigitNumbers() (ns FourDigitNumbers) {
	for _, k := range ks {
		ns = append(ns, k.ToFourDigitNumber())
	}
	return
}

func (ks FourDigitKanjis) ToString() string {
	str := ""
	for _, k := range ks {
		str = strings.Join(k.V, "") + k.E + str
	}
	return str
}

func (n FourDigitNumber) ToFourDigitKanji() (fourDigitKanji FourDigitKanji) {
	return FourDigitKanji{V: n.kanjiV(), E: n.kanjiE()}
}

func (n FourDigitNumber) kanjiE() string {
	for k, v := range LargePowerNumeralSymbols {
		if v == n.E {
			return k
		}
	}
	return ""
}

// TODO: リファクタ
// TODO: 一九〇四万みたいなフォーマットにも対応する
// TODO: どのようなフォーマットに変換するか選択できるようにする
func (n FourDigitNumber) kanjiV() (s []string) {
	v := n.V
	digits := []string{}
	for v > 0 {
		digits = append(digits, findArabicNumeralKanji(v%10))
		v /= 10
	}
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == "〇" {
			continue
		}
		if digits[i] != "一" {
			s = append(s, digits[i])
		}
		if i != 0 {
			s = append(s, findSmallPowerNumeralKanji(i))
		}
	}
	return
}

func (ns FourDigitNumbers) ToFourDigitKanjis() (ks FourDigitKanjis) {
	for _, n := range ns {
		ks = append(ks, n.ToFourDigitKanji())
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
