package kanjinumerals

import (
	"strings"
)

// splitNumeralSymbols 数詞で分割
func splitNumeralSymbols(s string) []string {
	return strings.Split(s, "")
}

// splitToFourDigitKanjis 漢数字を4桁ごとに分ける
func splitToFourDigitKanjis(kanjiNumerals string) (fourDigitKanjis FourDigitKanjis) {
	kanjiNumeralSymbols := splitNumeralSymbols(kanjiNumerals)
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

func KanjiToInt(s string) int {
	fourDigitKanjis := splitToFourDigitKanjis(s)
	fourDigitNumbers := fourDigitKanjis.ToFourDigitNumbers()
	return fourDigitNumbers.ToInt()
}

func IntToKanji(number int) string {
	// TODO
	return ""
}

// splitToFourDigitNumbers 数値を4桁ごとに分ける
func splitToFourDigitNumbers(arabicNumerals int) (fourDigitNumbers FourDigitNumbers) {
	e := 0
	for arabicNumerals > 0 {
		fourDigitNumbers = append(
			fourDigitNumbers,
			FourDigitNumber{
				V: arabicNumerals % 10000,
				E: e,
			},
		)
		arabicNumerals /= 10000
		e += 4
	}
	return
}
