package kanjinumerals

import (
	"errors"
	"strings"
)

func KanjiToInt(s string) (int, error) {
	kanjiNumeralSymbols := splitNumeralSymbols(s)
	if err := validateKanjis(kanjiNumeralSymbols); err != nil {
		return 0, err
	}
	fourDigitKanjis := splitToFourDigitKanjis(kanjiNumeralSymbols)
	fourDigitNumbers := fourDigitKanjis.ToFourDigitNumbers()
	return fourDigitNumbers.ToInt(), nil
}

func IntToKanji(number int) string {
	fourDigitNumbers := splitToFourDigitNumbers(number)
	fourDigitKanjis := fourDigitNumbers.ToFourDigitKanjis()
	return fourDigitKanjis.ToString()
}

func validateKanjis(kanjis []string) error {
	kanjiSymbols := getKanjiSymbols()
	for _, kanji := range kanjis {
		if !contains(kanjiSymbols, kanji) {
			return errors.New("一〜九、十百千、万億兆　以外の漢字が含まれています。")
		}
	}
	return nil
}

// splitNumeralSymbols 数詞で分割
func splitNumeralSymbols(s string) []string {
	return strings.Split(s, "")
}

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
