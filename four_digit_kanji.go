package kanjinumerals

import (
	"math"
	"strings"
)

type FourDigitKanji struct {
	V []string // 4桁の数値
	E string   // 10の冪乗
}
type FourDigitKanjis []FourDigitKanji

// splitToFourDigitKanjis 漢数字を4桁ごとに分ける
func splitToFourDigitKanjis(kanjiNumeralSymbols []string) (fourDigitKanjis FourDigitKanjis) {
	stuck := []string{}
	for _, v := range kanjiNumeralSymbols {
		if _, ok := LargePowerNumeralSymbols[v]; ok {
			fourDigitKanjis = append(fourDigitKanjis, FourDigitKanji{V: stuck, E: v})
			stuck = []string{}
		} else {
			stuck = append(stuck, v)
		}
	}
	if len(stuck) > 0 {
		fourDigitKanjis = append(fourDigitKanjis, FourDigitKanji{V: stuck, E: ""})
	}
	return fourDigitKanjis
}

// FourDigitNumberに変換
func (k FourDigitKanji) ToFourDigitNumber() (n FourDigitNumber) {
	n = FourDigitNumber{V: k.numberV(), E: 0}
	if e, ok := LargePowerNumeralSymbols[k.E]; ok {
		n.E = e
	}
	return
}

// 万億兆などの漢数字が含まれるか
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

// TODO: リファクタ
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
