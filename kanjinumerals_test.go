package kanjinumerals

import (
	"testing"
)

func TestKanjiToInt(t *testing.T) {
	t.Parallel()
	type args struct {
		kanjiNumeral string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "千二百三十四", args: args{kanjiNumeral: "千二百三十四"}, want: 1234},
		{name: "五〇六〇七八九", args: args{kanjiNumeral: "五〇六〇七八九"}, want: 5060789},
		{name: "十二兆三千四百二億三千四百五十万三千四百五十六", args: args{kanjiNumeral: "十二兆三千四百二億三千四百五十万三千四百五十六"}, want: 12340234503456},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KanjiToInt(tt.args.kanjiNumeral); got != tt.want {
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
