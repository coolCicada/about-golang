package main

import (
	"fmt"
	"sort"
)

type tuple[A, B any] struct {
	first A
	second B
}

func main() {
	ints := []int{ 2, 3, 4, 5, 0, -1, 2, 4 };
	fmt.Println(ints)
	sort.Ints(ints)
	fmt.Println(ints)

	strs := []string{ "a", "ab", "cd", "d", "b" }
	fmt.Println(strs)
	sort.Strings(strs)
	fmt.Println(strs)

	ta := []tuple[int, string]{ { 2, "a"}, { 1, "b" }, { 1, "a" } };
	fmt.Println(ta)
	sort.Slice(ta, func(i, j int) bool {
		if ta[i].first < ta[j].first {
			return true
		}
		if ta[i].first == ta[j].first {
			return ta[i].second < ta[j].second;
		}
		return false
	})
	fmt.Println(ta)
}