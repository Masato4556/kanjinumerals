package kanjinumerals

import (
	"math"
)

type FourDigitNumber struct {
	V int // 4桁の数値
	E int // 10の冪乗
}
type FourDigitNumbers []FourDigitNumber

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
