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
