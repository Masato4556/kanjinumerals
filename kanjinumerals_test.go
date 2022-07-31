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
		{name: "十二兆三千四百二億三千四百五十万三千四百五十六", args: args{s: "十二兆三千四百二億三千四百五十万三千四百五十六"}, want: 12340234503456},
		{name: "五〇六〇七八九", args: args{s: "五〇六〇七八九"}, want: 5060789}, // TODO：このケースも通るようにする
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
		want FourDigitKanjis
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
			want: FourDigitKanjis{
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
				t.Errorf("splitToFourDigit() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
