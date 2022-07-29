package kanjinumerals

import (
	"log"
	"math"
	"strings"
)

var LargeNumeralSymbols = map[string]int{"万": 4, "億": 8, "兆": 12}
var MediumNumeralSymbols = map[string]int{"十": 1, "百": 2, "千": 3}
var SmallNumeralSymbols = map[string]int{"〇": 0, "一": 1, "二": 2, "三": 3, "四": 4, "五": 5, "六": 6, "七": 7, "八": 8, "九": 9}

type FourDigitKanji struct {
	V []string // 4桁の数値
	E string   // 10の冪乗
}

func (k FourDigitKanji) ToFourDigitNumber() (n FourDigitNumber) {
	n = FourDigitNumber{V: k.numberV(), E: 0}
	if e, ok := LargeNumeralSymbols[k.E]; ok {
		// TODO: Vの値を計算する
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

type FourDigitNumber struct {
	V int // 4桁の数値
	E int // 10の冪乗
}

// splitNumeralSymbols 数詞で分割
func splitNumeralSymbols(s string) []string {
	return strings.Split(s, "")
}

// splitToFourDigit 漢数字を4桁ごとに分ける
func splitToFourDigit(kanjiNumeralSymbols []string) (fourDigitKanji []FourDigitKanji) {
	stuck := []string{}
	for _, v := range kanjiNumeralSymbols {
		if _, ok := LargeNumeralSymbols[v]; ok {
			fourDigitKanji = append(fourDigitKanji, FourDigitKanji{V: stuck, E: v})
			stuck = []string{}
		} else {
			stuck = append(stuck, v)
		}
	}
	if len(stuck) > 0 {
		fourDigitKanji = append(fourDigitKanji, FourDigitKanji{V: stuck, E: ""})
	}
	return fourDigitKanji
}

func KanjiToInt(s string) int {
	kanjiNumeralSymbols := splitNumeralSymbols(s)
	fourDigitKanji := splitToFourDigit(kanjiNumeralSymbols)
	log.Printf("%v", fourDigitKanji)
	return 0
}

func IntToKanji(number int) string {
    // TODO
    return ""
}
