package kanjinumerals

import (
	"reflect"
	"testing"
)

func TestFourDigitKanji_ToFourDigitNumber(t *testing.T) {
	type fields struct {
		V []string
		E string
	}
	tests := []struct {
		name   string
		fields fields
		wantN  FourDigitNumber
	}{
		{
			name:   "十二兆",
			fields: fields{V: []string{"十", "二"}, E: "兆"},
			wantN:  FourDigitNumber{V: 12, E: 12},
		},
		{
			name:   "百二億",
			fields: fields{V: []string{"百", "二"}, E: "億"},
			wantN:  FourDigitNumber{V: 102, E: 8},
		},
		{
			name:   "二千九百五十七万",
			fields: fields{V: []string{"二", "千", "九", "百", "五", "十", "七"}, E: "万"},
			wantN:  FourDigitNumber{V: 2957, E: 4},
		},
		{
			name:   "二千九百五十七万",
			fields: fields{V: []string{"六", "千", "五", "十"}, E: ""},
			wantN:  FourDigitNumber{V: 6050, E: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := FourDigitKanji{
				V: tt.fields.V,
				E: tt.fields.E,
			}
			if gotN := k.ToFourDigitNumber(); !reflect.DeepEqual(gotN, tt.wantN) {
				t.Errorf("FourDigitKanji.ToFourDigitNumber() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestFourDigitKanji_numberV(t *testing.T) {
	type fields struct {
		V []string
	}
	tests := []struct {
		name   string
		fields fields
		wantNv int
	}{
		{
			name:   "",
			fields: fields{V: []string{"十", "二"}},
			wantNv: 12,
		},
		{
			name:   "",
			fields: fields{V: []string{"百", "二"}},
			wantNv: 102,
		},
		{
			name:   "",
			fields: fields{V: []string{"二", "千", "九", "百", "五", "十", "七"}},
			wantNv: 2957,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := FourDigitKanji{
				V: tt.fields.V,
			}
			if gotNv := k.numberV(); gotNv != tt.wantNv {
				t.Errorf("FourDigitKanji.numberV() = %v, want %v", gotNv, tt.wantNv)
			}
		})
	}
}

func TestFourDigitKanjis_ToFourDigitNumbers(t *testing.T) {
	tests := []struct {
		name   string
		ks     FourDigitKanjis
		wantNs FourDigitNumbers
	}{
		{
			name: "",
			ks: FourDigitKanjis{
				{V: []string{"十", "二"}, E: "兆"},
				{V: []string{"百", "二"}, E: "億"},
				{V: []string{"二", "千", "九", "百", "五", "十", "七"}, E: "万"},
				{V: []string{"六", "千", "五", "十"}, E: ""},
			},
			wantNs: FourDigitNumbers{
				{V: 12, E: 12},
				{V: 102, E: 8},
				{V: 2957, E: 4},
				{V: 6050, E: 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNs := tt.ks.ToFourDigitNumbers(); !reflect.DeepEqual(gotNs, tt.wantNs) {
				t.Errorf("FourDigitKanjis.ToFourDigitNumbers() = %v, want %v", gotNs, tt.wantNs)
			}
		})
	}
}

func TestFourDigitNumbers_ToInt(t *testing.T) {
	tests := []struct {
		name string
		ns   FourDigitNumbers
		want int
	}{
		{
			name: "",
			ns: FourDigitNumbers{
				{V: 12, E: 12},
				{V: 102, E: 8},
				{V: 2957, E: 4},
				{V: 6050, E: 0},
			},
			want: 12010229576050,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ns.ToInt(); got != tt.want {
				t.Errorf("FourDigitNumbers.ToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFourDigitKanji_IncludeSmallPowerNumeralSymbols(t *testing.T) {
	type fields struct {
		V []string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{name: "一", fields: fields{V: []string{"一"}}, want: false},
		{name: "十", fields: fields{V: []string{"十"}}, want: true},
		{name: "二百五", fields: fields{V: []string{"二", "百", "五"}}, want: true},
		{name: "二〇五", fields: fields{V: []string{"二", "〇", "五"}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := FourDigitKanji{
				V: tt.fields.V,
			}
			if got := k.IncludeSmallPowerNumeralSymbols(); got != tt.want {
				t.Errorf("FourDigitKanji.IncludeSmallPowerNumeralSymbols() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFourDigitKanji_vToNumberWithPowers(t *testing.T) {
	type fields struct {
		V []string
	}
	tests := []struct {
		name   string
		fields fields
		wantNv int
	}{
		{name: "一", fields: fields{V: []string{"一"}}, wantNv: 1},
		{name: "十", fields: fields{V: []string{"十"}}, wantNv: 10},
		{name: "二百五", fields: fields{V: []string{"二", "百", "五"}}, wantNv: 205},
		{name: "三千六百十五", fields: fields{V: []string{"三", "千", "六", "百", "十", "五"}}, wantNv: 3615},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := FourDigitKanji{
				V: tt.fields.V,
			}
			if gotNv := k.vToNumberWithPowers(); gotNv != tt.wantNv {
				t.Errorf("FourDigitKanji.vToNumberWithPowers() = %v, want %v", gotNv, tt.wantNv)
			}
		})
	}
}

func TestFourDigitKanji_vToNumberWithoutPowers(t *testing.T) {
	type fields struct {
		V []string
		E string
	}
	tests := []struct {
		name   string
		fields fields
		wantNv int
	}{
		{name: "二〇五", fields: fields{V: []string{"二", "〇", "五"}}, wantNv: 205},
		{name: "三五六七", fields: fields{V: []string{"三", "五", "六", "七"}}, wantNv: 3567},
		{name: "一二三四五", fields: fields{V: []string{"一", "二", "三", "四", "五"}}, wantNv: 12345},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := FourDigitKanji{
				V: tt.fields.V,
				E: tt.fields.E,
			}
			if gotNv := k.vToNumberWithoutPowers(); gotNv != tt.wantNv {
				t.Errorf("FourDigitKanji.vToNumberWithoutPowers() = %v, want %v", gotNv, tt.wantNv)
			}
		})
	}
}

func TestFourDigitNumber_kanjiV(t *testing.T) {
	type fields struct {
		V int
	}
	tests := []struct {
		name   string
		fields fields
		wantS  []string
	}{
		{name: "1925", fields: fields{V: 1925}, wantS: []string{"千", "九", "百", "二", "十", "五"}},
		{name: "806", fields: fields{V: 806}, wantS: []string{"八", "百", "六"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := FourDigitNumber{
				V: tt.fields.V,
			}
			if gotS := n.kanjiV(); !reflect.DeepEqual(gotS, tt.wantS) {
				t.Errorf("FourDigitNumber.kanjiV() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func TestFourDigitNumber_kanjiE(t *testing.T) {
	type fields struct {
		E int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "0", fields: fields{E: 0}, want: ""},
		{name: "4", fields: fields{E: 4}, want: "万"},
		{name: "8", fields: fields{E: 8}, want: "億"},
		{name: "12", fields: fields{E: 12}, want: "兆"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := FourDigitNumber{
				E: tt.fields.E,
			}
			if got := n.kanjiE(); got != tt.want {
				t.Errorf("FourDigitNumber.kanjiE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFourDigitNumber_ToFourDigitKanji(t *testing.T) {
	type fields struct {
		V int
		E int
	}
	tests := []struct {
		name               string
		fields             fields
		wantFourDigitKanji FourDigitKanji
	}{
		{name: "12", fields: fields{V: 1923, E: 12}, wantFourDigitKanji: FourDigitKanji{V: []string{"千", "九", "百", "二", "十", "三"}, E: "兆"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := FourDigitNumber{
				V: tt.fields.V,
				E: tt.fields.E,
			}
			if gotFourDigitKanji := n.ToFourDigitKanji(); !reflect.DeepEqual(gotFourDigitKanji, tt.wantFourDigitKanji) {
				t.Errorf("FourDigitNumber.ToFourDigitKanji() = %v, want %v", gotFourDigitKanji, tt.wantFourDigitKanji)
			}
		})
	}
}
