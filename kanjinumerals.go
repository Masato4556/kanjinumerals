package kanjinumerals

import (
	"errors"
	"math/big"
	"strings"
)

func KanjiToInt(s string) (*big.Int, error) {
	kanjiNumeralSymbols := splitNumeralSymbols(s)
	if err := validateKanjis(kanjiNumeralSymbols); err != nil {
		return genBigInt0(), err
	}
	fourDigitKanjis := splitToFourDigitKanjis(kanjiNumeralSymbols)
	fourDigitNumbers := fourDigitKanjis.ToFourDigitNumbers()
	return fourDigitNumbers.ToInt(), nil
}

func IntToKanji(number *big.Int) (string, error) {
	fourDigitNumbers := splitToFourDigitNumbers(number)
	fourDigitKanjis := fourDigitNumbers.ToFourDigitKanjis()
	return fourDigitKanjis.ToString(), nil
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
