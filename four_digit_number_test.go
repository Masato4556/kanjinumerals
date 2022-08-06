package kanjinumerals

import (
	"reflect"
	"testing"
)

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
