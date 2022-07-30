package kanjinumerals

import (
	"log"
	"strings"
)

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

// // 漢数字
// type KanjiNumeralSymbol interface {
// 	calculate(int) int
// }

// // 一 ~ 九
// type KanjiNumeralSmallSymbol struct {
// 	value int
// }

// func (s KanjiNumeralSmallSymbol) calculate(int) int {
// 	return 0
// }

// // 十、百、千
// type KanjiNumeralMediumSymbol struct {
// 	value int
// }

// func (s KanjiNumeralMediumSymbol) calculate(int) int {
// 	return 0
// }

// // 万、億、兆など
// type KanjiNumeralBigSymbol struct {
// 	value int
// }

// func (s KanjiNumeralBigSymbol) calculate(int) int {
// 	return 0
// }

/*
大まかな変換の流れ

1: 万、億、兆ごとに分割する
十二兆三千四百二億三千四百五十万三千四百五十六　＝＞　[十二兆, 三千四百二億, 三千四百五十万, 三千四百五十六]

2: 万、億、兆の部分と、それ以外の部分で分割する
[十二兆, 三千四百二億, 三千四百五十万, 三千四百五十六] => [(十二, 兆), (三千四百二, 億), (三千四百五十, 万), (三千四百五十六, )]

3: 万、億、兆以外の部分はそのままintに変換、億、兆の部分は10の次数に変換
[(十二, 兆), (三千四百二, 億), (三千四百五十, 万), (三千四百五十六)] => [(12, 12), (3402, 8), (3450, 4), (3456, 0)]

4 最終的な数値を出す。
[(12, 12), (3402, 8), (3450, 4), (3456, 0)] => 12*(10^12) + 3402*(10^8) + 3450*(10^4) + 3456*(10^0)


上記の3の時の計算方法
3-1: 十、百、千の部分で分ける
三千四百五十 => [三千, 四百,五十]
3-2: 十、百、千の部分と、それ以外の部分で分割する
[三千, 四百,五十] => [(三, 千), (四, 百), (五, 十)]
3-3: 十、百、千以外の部分はそのままintに変換、十、百、千の部分は10の次数に変換
[(三, 千), (四, 百), (五, 十)] => [(3, 3), (4, 2), (5, 1)]
4 最終的な数値を出す。
 [(3, 3), (4, 2), (5, 1)] =>  3*(10^3) + 4*(10^2)) + 5*(10^1)

（二〇二二, ) => (2022, 1)

*/

// var kanjiNumeralSymbols = map[string]KanjiNumeralSymbol{
// 	"〇": KanjiNumeralSmallSymbol{value: 0},
// 	"一": KanjiNumeralSmallSymbol{value: 1},
// 	"二": KanjiNumeralSmallSymbol{value: 2},
// 	"三": KanjiNumeralSmallSymbol{value: 3},
// 	"四": KanjiNumeralSmallSymbol{value: 4},
// 	"五": KanjiNumeralSmallSymbol{value: 5},
// 	"六": KanjiNumeralSmallSymbol{value: 6},
// 	"七": KanjiNumeralSmallSymbol{value: 7},
// 	"八": KanjiNumeralSmallSymbol{value: 8},
// 	"九": KanjiNumeralSmallSymbol{value: 9},

// 	"十": KanjiNumeralMediumSymbol{value: 10},
// 	"百": KanjiNumeralMediumSymbol{value: 100},
// 	"千": KanjiNumeralMediumSymbol{value: 1000},

// 	"万": KanjiNumeralBigSymbol{value: 10000},
// 	"億": KanjiNumeralBigSymbol{value: int(math.Pow(10, 8))},

// 	// "零": KanjiNumeralSmallSymbol{value: 0},
// 	// "壱": KanjiNumeralSmallSymbol{value: 1},
// 	// "弐": KanjiNumeralSmallSymbol{value: 2},
// 	// "参": KanjiNumeralSmallSymbol{value: 3},
// 	// "壹": KanjiNumeralSmallSymbol{value: 1},
// 	// "貳": KanjiNumeralSmallSymbol{value: 2},
// 	// "叁": KanjiNumeralSmallSymbol{value: 3},
// 	// "肆": KanjiNumeralSmallSymbol{value: 4},
// 	// "伍": KanjiNumeralSmallSymbol{value: 5},
// 	// "陸": KanjiNumeralSmallSymbol{value: 6},
// 	// "柒": KanjiNumeralSmallSymbol{value: 7},
// 	// "捌": KanjiNumeralSmallSymbol{value: 8},
// 	// "玖": KanjiNumeralSmallSymbol{value: 9},

// 	// "０": KanjiNumeralSmallSymbol{value: 0},
// 	// "１": KanjiNumeralSmallSymbol{value: 1},
// 	// "２": KanjiNumeralSmallSymbol{value: 2},
// 	// "３": KanjiNumeralSmallSymbol{value: 3},
// 	// "４": KanjiNumeralSmallSymbol{value: 4},
// 	// "５": KanjiNumeralSmallSymbol{value: 5},
// 	// "６": KanjiNumeralSmallSymbol{value: 6},
// 	// "７": KanjiNumeralSmallSymbol{value: 7},
// 	// "８": KanjiNumeralSmallSymbol{value: 8},
// 	// "９": KanjiNumeralSmallSymbol{value: 9},

// 	// "拾": KanjiNumeralMediumSymbol{value: 1}0,
// 	// "陌": KanjiNumeralMediumSymbol{value: 1}00,
// 	// "佰": KanjiNumeralMediumSymbol{value: 1}00,
// 	// "阡": KanjiNumeralMediumSymbol{value: 1}000,
// 	// "仟": KanjiNumeralMediumSymbol{value: 1}000,

// 	// "萬": KanjiNumeralSmallSymbol{value: 10000},
// }

// https://ja.wikipedia.org/wiki/%E5%91%BD%E6%95%B0%E6%B3%95
