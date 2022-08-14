package kanjinumerals

import "math/big"

func contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

func merge[T comparable](m ...map[string]T) map[string]T {
	ans := make(map[string]T, 0)
	for _, c := range m {
		for k, v := range c {
			ans[k] = v
		}
	}
	return ans
}

func keys[T comparable](m map[string]T) []string {
	ans := make([]string, 0)
	for key := range m {
		ans = append(ans, key)
	}
	return ans
}

//===== big.Int =====

var bigInt0 = big.NewInt(0)
var bigInt10 = big.NewInt(10)

func genBigInt0() *big.Int {
	return new(big.Int).Set(bigInt0)
}

func genBigInt10() *big.Int {
	return new(big.Int).Set(bigInt10)
}

func cmpZero(n *big.Int) int {
	return n.Cmp(genBigInt0())
}
