package kanjinumerals

import (
	"math/big"
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
	n = FourDigitNumber{V: k.numberV(), E: genBigInt0()}
	if e, ok := LargePowerNumeralSymbols[k.E]; ok {
		n.E = big.NewInt(e)
	}
	return
}

// 万億兆などの漢数字が含まれるか
func (k FourDigitKanji) IncludeSmallPowerNumeralSymbols() bool {
	for symbol := range SmallPowerNumeralSymbols {
		if contains(k.V, symbol) {
			return true
		}
	}
	return false
}

func (k FourDigitKanji) numberV() (nv *big.Int) {
	if k.IncludeSmallPowerNumeralSymbols() {
		return k.vToNumberWithPowers()
	}

	return k.vToNumberWithoutPowers()
}

// TODO: リファクタ
func (k FourDigitKanji) vToNumberWithPowers() *big.Int {
	temp := genBigInt0()
	digits := []*big.Int{}
	nv := genBigInt0()
	for _, v := range k.V {
		if mns, ok := SmallPowerNumeralSymbols[v]; ok {
			if cmpZero(temp) == 0 {
				temp.Set(big.NewInt(1))
			}
			temp.Mul(temp, new(big.Int).Exp(genBigInt10(), big.NewInt(mns), nil))
			digits = append(digits, new(big.Int).Set(temp))
			temp.Set(genBigInt0())
		}
		if sns, ok := ArabicNumeralSymbols[v]; ok {
			temp.Add(temp, big.NewInt(sns))
		}
	}
	if cmpZero(temp) != 0 {
		digits = append(digits, temp)
	}
	for _, v := range digits {
		nv.Add(nv, v)
	}
	return nv
}

func (k FourDigitKanji) vToNumberWithoutPowers() *big.Int {
	nv := genBigInt0()
	for i, v := range k.V {
		if i > 0 {
			nv.Mul(nv, genBigInt10())
		}
		if sns, ok := ArabicNumeralSymbols[v]; ok {
			nv.Add(nv, big.NewInt(sns))
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
