package kanjinumerals

import (
	"log"
	"strings"
)

var LargeNumeralSymbols = []string{"万", "億", "兆"}

type FourDigitKanji struct {
	V []string // 4桁の数値
	E string   // 10の冪乗
}

// splitNumeralSymbols 数詞で分割
func splitNumeralSymbols(s string) []string {
	return strings.Split(s, "")
}

// splitToFourDigit 漢数字を4桁ごとに分ける
func splitToFourDigit(kanjiNumeralSymbols []string) (fourDigitKanji []FourDigitKanji) {
	stuck := []string{}
	for _, v := range kanjiNumeralSymbols {
		if contains(LargeNumeralSymbols, v) {
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
