package kanjinumerals

import (
	"reflect"
	"testing"
)

func TestKanjiToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{name: "千二百三十四", args: args{s: "千二百三十四"}, want: 1234, wantErr: false},
		{name: "十二兆三千四百二億三千四百五十万三千四百五十六", args: args{s: "十二兆三千四百二億三千四百五十万三千四百五十六"}, want: 12340234503456, wantErr: false},
		{name: "五〇六〇七八九", args: args{s: "五〇六〇七八九"}, want: 5060789, wantErr: false},
		{name: "漢数字以外", args: args{s: "五〇漢字"}, want: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := KanjiToInt(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("KanjiToInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
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

func Test_splitToFourDigitKanjis(t *testing.T) {
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
			args: args{kanjiNumeral: []string{
				"十", "二", "兆",
				"三", "千", "四", "百", "二", "億",
				"三", "千", "四", "百", "五", "十", "万",
				"三", "千", "四", "百", "五", "十", "六",
			}},
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
			if got := splitToFourDigitKanjis(tt.args.kanjiNumeral); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitToFourDigitKanjis() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func Test_splitToFourDigitNumbers(t *testing.T) {
	type args struct {
		arabicNumerals int
	}
	tests := []struct {
		name                 string
		args                 args
		wantFourDigitNumbers FourDigitNumbers
	}{
		{name: "123", args: args{arabicNumerals: 123}, wantFourDigitNumbers: FourDigitNumbers{{V: 123, E: 0}}},
		{
			name: "90347389462590", args: args{arabicNumerals: 90347389462590},
			wantFourDigitNumbers: FourDigitNumbers{
				{V: 2590, E: 0},
				{V: 8946, E: 4},
				{V: 3473, E: 8},
				{V: 90, E: 12},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFourDigitNumbers := splitToFourDigitNumbers(tt.args.arabicNumerals); !reflect.DeepEqual(gotFourDigitNumbers, tt.wantFourDigitNumbers) {
				t.Errorf("splitToFourDigitNumbers() = %v, want %v", gotFourDigitNumbers, tt.wantFourDigitNumbers)
			}
		})
	}
}
