# kanjinumerals
Converting Japanese Kanji numerals and numbers.

## Installation

```
go get -u github.com/Masato4556/kanjinumerals
```

## Usage

```go
package main

import "github.com/Masato4556/kanjinumerals"

func main() {
	println(kanjinumerals.IntToKanji(980032))
	println(kanjinumerals.KanjiToInt("百万五二"))
}
```