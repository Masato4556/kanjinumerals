package kanjinumerals

import (
	"math/big"
)

type FourDigitNumber struct {
	V *big.Int // 4桁の数値
	E *big.Int // 10の冪乗
}
type FourDigitNumbers []FourDigitNumber

// splitToFourDigitNumbers 数値を4桁ごとに分ける
func splitToFourDigitNumbers(arabicNumerals *big.Int) (fourDigitNumbers FourDigitNumbers) {
	e := genBigInt0()
	var m *big.Int
	for cmpZero(arabicNumerals) > 0 {
		arabicNumerals, m = arabicNumerals.DivMod(
			arabicNumerals,
			new(big.Int).Exp(genBigInt10(), big.NewInt(4), nil),
			new(big.Int),
		)
		if cmpZero(m) != 0 {
			fourDigitNumbers = append(
				fourDigitNumbers,
				FourDigitNumber{
					V: new(big.Int).Set(m),
					E: new(big.Int).Set(e),
				},
			)
		}
		e.Add(e, big.NewInt(4))
	}
	return
}

func (n FourDigitNumber) ToFourDigitKanji() (fourDigitKanji FourDigitKanji) {
	return FourDigitKanji{V: n.kanjiV(), E: n.kanjiE()}
}

func (n FourDigitNumber) kanjiE() string {
	for k, v := range LargePowerNumeralSymbols {
		if big.NewInt(v).Cmp(n.E) == 0 {
			return k
		}
	}
	return ""
}

// TODO: リファクタ
// TODO: 一九〇四万みたいなフォーマットにも対応する
// TODO: どのようなフォーマットに変換するか選択できるようにする
func (n FourDigitNumber) kanjiV() (s []string) {
	if n.V == nil {
		return []string{}
	}

	v := n.V
	var m *big.Int
	ten := genBigInt10()
	digits := []string{}
	for cmpZero(v) > 0 {
		v, m = new(big.Int).DivMod(v, ten, new(big.Int)) // TODO: DivModの使い方をちゃんと理解する
		digits = append(digits, findArabicNumeralKanji(m.Int64()))
	}
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == "〇" { // TODO:　ロジックの最適化
			continue
		}
		if digits[i] != "一" || i == 0 { // TODO:　ロジックの最適化
			s = append(s, digits[i])
		}
		if i != 0 {
			s = append(s, findSmallPowerNumeralKanji(int64(i)))
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
func (ns FourDigitNumbers) ToInt() *big.Int {

	number := genBigInt0()
	ten := genBigInt10()
	var e big.Int
	for _, n := range ns {
		e.Exp(ten, n.E, nil)
		number.Add(number, new(big.Int).Mul(n.V, &e))
	}
	return number
}
