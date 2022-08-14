package kanjinumerals

import (
	"math/big"
	"reflect"
	"testing"
)

var tenExp48 = new(big.Int).Exp(genBigInt10(), big.NewInt(48), nil)

func TestKanjiToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr bool
	}{
		{name: "千二百三十四", args: args{s: "千二百三十四"}, want: big.NewInt(1234), wantErr: false},
		{name: "十二兆三千四百二億三千四百五十万三千四百五十六", args: args{s: "十二兆三千四百二億三千四百五十万三千四百五十六"}, want: big.NewInt(12340234503456), wantErr: false},
		{name: "五〇六〇七八九", args: args{s: "五〇六〇七八九"}, want: big.NewInt(5060789), wantErr: false},
		{name: "一極", args: args{s: "一極"}, want: tenExp48, wantErr: false},
		{name: "漢数字以外", args: args{s: "五〇漢字"}, want: genBigInt0(), wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := KanjiToInt(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("KanjiToInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KanjiToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntToKanji(t *testing.T) {
	type args struct {
		number *big.Int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "1234", args: args{number: big.NewInt(1234)}, want: "千二百三十四", wantErr: false},
		{name: "12340234503456", args: args{number: big.NewInt(12340234503456)}, want: "十二兆三千四百二億三千四百五十万三千四百五十六", wantErr: false},
		{name: "1極", args: args{number: tenExp48}, want: "一極", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IntToKanji(tt.args.number)
			if (err != nil) != tt.wantErr {
				t.Errorf("IntToKanji() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IntToKanji() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateKanjis(t *testing.T) {
	type args struct {
		kanjis []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "空", args: args{kanjis: []string{}}, wantErr: false},
		{name: "正常", args: args{kanjis: []string{"十", "二", "億", "三", "千", "四", "百", "五", "十", "六", "万", "七", "千", "八", "百", "九", "十"}}, wantErr: false},
		{name: "漢数字以外", args: args{kanjis: []string{"漢"}}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateKanjis(tt.args.kanjis); (err != nil) != tt.wantErr {
				t.Errorf("validateKanjis() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_splitNumeralSymbols(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "", args: args{s: "十二億三千四百五十六万七千八百九十"}, want: []string{"十", "二", "億", "三", "千", "四", "百", "五", "十", "六", "万", "七", "千", "八", "百", "九", "十"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitNumeralSymbols(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitNumeralSymbols() = %v, want %v", got, tt.want)
			}
		})
	}
}
