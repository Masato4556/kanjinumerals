package kanjinumerals

var LargePowerNumeralSymbols = map[string]int{"万": 4, "億": 8, "兆": 12}
var SmallPowerNumeralSymbols = map[string]int{"十": 1, "百": 2, "千": 3}
var ArabicNumeralSymbols = map[string]int{"〇": 0, "一": 1, "二": 2, "三": 3, "四": 4, "五": 5, "六": 6, "七": 7, "八": 8, "九": 9}

func findSmallPowerNumeralKanji(n int) string {
	for k, v := range SmallPowerNumeralSymbols {
		if v == n {
			return k
		}
	}
	return ""
}

func findArabicNumeralKanji(n int) string {
	for k, v := range ArabicNumeralSymbols {
		if v == n {
			return k
		}
	}
	return ""
}

func getKanjiSymbols() []string {
	return keys(merge(LargePowerNumeralSymbols, SmallPowerNumeralSymbols, ArabicNumeralSymbols))
}
