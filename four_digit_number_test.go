package kanjinumerals

import (
	"math/big"
	"reflect"
	"testing"
)

// func TestFourDigitNumber_ToFourDigitKanji(t *testing.T) {
// 	type fields struct {
// 		V int
// 		E int
// 	}
// 	tests := []struct {
// 		name               string
// 		fields             fields
// 		wantFourDigitKanji FourDigitKanji
// 	}{
// 		{name: "12", fields: fields{V: 1923, E: 12}, wantFourDigitKanji: FourDigitKanji{V: []string{"千", "九", "百", "二", "十", "三"}, E: "兆"}},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			n := FourDigitNumber{
// 				V: tt.fields.V,
// 				E: tt.fields.E,
// 			}
// 			if gotFourDigitKanji := n.ToFourDigitKanji(); !reflect.DeepEqual(gotFourDigitKanji, tt.wantFourDigitKanji) {
// 				t.Errorf("FourDigitNumber.ToFourDigitKanji() = %v, want %v", gotFourDigitKanji, tt.wantFourDigitKanji)
// 			}
// 		})
// 	}
// }

func TestFourDigitNumbers_ToInt(t *testing.T) {
	tests := []struct {
		name string
		ns   FourDigitNumbers
		want *big.Int
	}{
		{
			name: "",
			ns: FourDigitNumbers{
				{V: big.NewInt(12), E: big.NewInt(12)},
				{V: big.NewInt(102), E: big.NewInt(8)},
				{V: big.NewInt(2957), E: big.NewInt(4)},
				{V: big.NewInt(6050), E: big.NewInt(0)},
			},
			want: big.NewInt(12010229576050),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ns.ToInt(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FourDigitNumbers.ToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFourDigitNumber_kanjiV(t *testing.T) {
	type fields struct {
		V *big.Int
		E *big.Int
	}
	tests := []struct {
		name   string
		fields fields
		wantS  []string
	}{
		{name: "1925", fields: fields{V: big.NewInt(1925)}, wantS: []string{"千", "九", "百", "二", "十", "五"}},
		{name: "806", fields: fields{V: big.NewInt(806)}, wantS: []string{"八", "百", "六"}},
		{name: "nil", fields: fields{V: nil}, wantS: []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := FourDigitNumber{
				V: tt.fields.V,
				E: tt.fields.E,
			}
			if gotS := n.kanjiV(); !reflect.DeepEqual(gotS, tt.wantS) {
				t.Errorf("FourDigitNumber.kanjiV() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func TestFourDigitNumber_kanjiE(t *testing.T) {
	type fields struct {
		V *big.Int
		E *big.Int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "0", fields: fields{E: big.NewInt(0)}, want: ""},
		{name: "4", fields: fields{E: big.NewInt(4)}, want: "万"},
		{name: "8", fields: fields{E: big.NewInt(8)}, want: "億"},
		{name: "12", fields: fields{E: big.NewInt(12)}, want: "兆"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := FourDigitNumber{
				V: tt.fields.V,
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
		V *big.Int
		E *big.Int
	}
	tests := []struct {
		name               string
		fields             fields
		wantFourDigitKanji FourDigitKanji
	}{
		{name: "12", fields: fields{V: big.NewInt(1923), E: big.NewInt(12)}, wantFourDigitKanji: FourDigitKanji{V: []string{"千", "九", "百", "二", "十", "三"}, E: "兆"}},
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

func Test_splitToFourDigitNumbers(t *testing.T) {
	type args struct {
		arabicNumerals *big.Int
	}
	tests := []struct {
		name                 string
		args                 args
		wantFourDigitNumbers FourDigitNumbers
	}{
		{name: "", args: args{arabicNumerals: big.NewInt(1234567890)},
			wantFourDigitNumbers: FourDigitNumbers{
				{V: big.NewInt(7890), E: big.NewInt(0)},
				{V: big.NewInt(3456), E: big.NewInt(4)},
				{V: big.NewInt(12), E: big.NewInt(8)},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFourDigitNumbers := splitToFourDigitNumbers(tt.args.arabicNumerals); !reflect.DeepEqual(gotFourDigitNumbers, tt.wantFourDigitNumbers) {
				t.Errorf("splitToFourDigitNumbers() = %v, want %v", gotFourDigitNumbers, tt.wantFourDigitNumbers)
			}
		})
	}
}
