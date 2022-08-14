package kanjinumerals

import (
	"math/big"
	"reflect"
	"testing"
)

func TestFourDigitKanji_vToNumberWithPowers(t *testing.T) {
	type fields struct {
		V []string
		E string
	}
	tests := []struct {
		name   string
		fields fields
		wantNv *big.Int
	}{
		{name: "一", fields: fields{V: []string{"一"}}, wantNv: big.NewInt(1)},
		{name: "十", fields: fields{V: []string{"十"}}, wantNv: genBigInt10()},
		{name: "二百五", fields: fields{V: []string{"二", "百", "五"}}, wantNv: big.NewInt(205)},
		{name: "三千六百十五", fields: fields{V: []string{"三", "千", "六", "百", "十", "五"}}, wantNv: big.NewInt(3615)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := FourDigitKanji{
				V: tt.fields.V,
				E: tt.fields.E,
			}
			if gotNv := k.vToNumberWithPowers(); !reflect.DeepEqual(gotNv, tt.wantNv) {
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
		wantNv *big.Int
	}{
		{name: "二〇五", fields: fields{V: []string{"二", "〇", "五"}}, wantNv: big.NewInt(205)},
		{name: "三五六七", fields: fields{V: []string{"三", "五", "六", "七"}}, wantNv: big.NewInt(3567)},
		{name: "一二三四五", fields: fields{V: []string{"一", "二", "三", "四", "五"}}, wantNv: big.NewInt(12345)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := FourDigitKanji{
				V: tt.fields.V,
				E: tt.fields.E,
			}
			if gotNv := k.vToNumberWithoutPowers(); !reflect.DeepEqual(gotNv, tt.wantNv) {
				t.Errorf("FourDigitKanji.vToNumberWithoutPowers() = %v, want %v", gotNv, tt.wantNv)
			}
		})
	}
}

func TestFourDigitKanji_IncludeSmallPowerNumeralSymbols(t *testing.T) {
	type fields struct {
		V []string
		E string
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
				E: tt.fields.E,
			}
			if got := k.IncludeSmallPowerNumeralSymbols(); got != tt.want {
				t.Errorf("FourDigitKanji.IncludeSmallPowerNumeralSymbols() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFourDigitKanji_numberV(t *testing.T) {
	type fields struct {
		V []string
		E string
	}
	tests := []struct {
		name   string
		fields fields
		wantNv *big.Int
	}{
		{
			name:   "十二",
			fields: fields{V: []string{"十", "二"}},
			wantNv: big.NewInt(12),
		},
		{
			name:   "一二",
			fields: fields{V: []string{"一", "二"}},
			wantNv: big.NewInt(12),
		},
		{
			name:   "二千九百五十七",
			fields: fields{V: []string{"二", "千", "九", "百", "五", "十", "七"}},
			wantNv: big.NewInt(2957),
		},
		{
			name:   "二九五七",
			fields: fields{V: []string{"二", "九", "五", "七"}},
			wantNv: big.NewInt(2957),
		},
		// TODO: 五〇〇万みたいなケースにも対応する
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := FourDigitKanji{
				V: tt.fields.V,
				E: tt.fields.E,
			}
			if gotNv := k.numberV(); !reflect.DeepEqual(gotNv, tt.wantNv) {
				t.Errorf("FourDigitKanji.numberV() = %v, want %v", gotNv, tt.wantNv)
			}
		})
	}
}

func Test_splitToFourDigitKanjis(t *testing.T) {
	type args struct {
		kanjiNumeralSymbols []string
	}
	tests := []struct {
		name                string
		args                args
		wantFourDigitKanjis FourDigitKanjis
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFourDigitKanjis := splitToFourDigitKanjis(tt.args.kanjiNumeralSymbols); !reflect.DeepEqual(gotFourDigitKanjis, tt.wantFourDigitKanjis) {
				t.Errorf("splitToFourDigitKanjis() = %v, want %v", gotFourDigitKanjis, tt.wantFourDigitKanjis)
			}
		})
	}
}

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
			wantN:  FourDigitNumber{V: big.NewInt(12), E: big.NewInt(12)},
		},
		{
			name:   "百二億",
			fields: fields{V: []string{"百", "二"}, E: "億"},
			wantN:  FourDigitNumber{V: big.NewInt(102), E: big.NewInt(8)},
		},
		{
			name:   "二千九百五十七万",
			fields: fields{V: []string{"二", "千", "九", "百", "五", "十", "七"}, E: "万"},
			wantN:  FourDigitNumber{V: big.NewInt(2957), E: big.NewInt(4)},
		},
		{
			name:   "二千九百五十七万",
			fields: fields{V: []string{"六", "千", "五", "十"}, E: ""},
			wantN:  FourDigitNumber{V: big.NewInt(6050), E: genBigInt0()},
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
				{V: big.NewInt(12), E: big.NewInt(12)},
				{V: big.NewInt(102), E: big.NewInt(8)},
				{V: big.NewInt(2957), E: big.NewInt(4)},
				{V: big.NewInt(6050), E: genBigInt0()},
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
