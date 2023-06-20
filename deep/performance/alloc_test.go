package performance

import "testing"

func BenchmarkNoPre(b *testing.B) {
	for n := 0; n < b.N; n ++ {
		NoPreAlloc(10000)
	}
}

func BenchmarkPre(b *testing.B) {
	for n := 0; n < b.N; n ++ {
		PreAlloc(10000)
	}
}