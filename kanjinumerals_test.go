package kanjinumerals

import (
	"reflect"
	"testing"
)

func TestKanjiToInt(t *testing.T) {
	t.Parallel()
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "千二百三十四", args: args{s: "千二百三十四"}, want: 1234},
		{name: "五〇六〇七八九", args: args{s: "五〇六〇七八九"}, want: 5060789},
		{name: "十二兆三千四百二億三千四百五十万三千四百五十六", args: args{s: "十二兆三千四百二億三千四百五十万三千四百五十六"}, want: 12340234503456},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KanjiToInt(tt.args.s); got != tt.want {
				t.Errorf("KanjiToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntToKanji(t *testing.T) {
	t.Parallel()
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1234", args: args{number: 1234}, want: "千二百三十四"},
		{name: "5060789", args: args{number: 5060789}, want: "五〇六〇七八九"},
		{name: "12340234503456", args: args{number: 12340234503456}, want: "十二兆三千四百二億三千四百五十万三千四百五十六"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntToKanji(tt.args.number); got != tt.want {
				t.Errorf("IntToKanji() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitToFourDigit(t *testing.T) {
	type args struct {
		kanjiNumeral []string
	}
	tests := []struct {
		name string
		args args
		want []FourDigitKanji
	}{
		{
			name: "十二兆三千四百二億三千四百五十万三千四百五十六",
			args: args{
				kanjiNumeral: []string{
					"十", "二", "兆",
					"三", "千", "四", "百", "二", "億",
					"三", "千", "四", "百", "五", "十", "万",
					"三", "千", "四", "百", "五", "十", "六",
				},
			},
			want: []FourDigitKanji{
				{V: []string{"十", "二"}, E: "兆"},
				{V: []string{"三", "千", "四", "百", "二"}, E: "億"},
				{V: []string{"三", "千", "四", "百", "五", "十"}, E: "万"},
				{V: []string{"三", "千", "四", "百", "五", "十", "六"}, E: ""},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitToFourDigit(tt.args.kanjiNumeral); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitToFourDigit() = %v, want %v", got, tt.want)
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
