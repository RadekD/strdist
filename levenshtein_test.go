package strdist_test

import (
	"testing"

	"github.com/RadekD/strdist"
)

var testCases = []struct {
	source   string
	target   string
	distance int
}{
	{"book", "back", 2},
	{"brook", "back", 3},
	{"berk", "back", 2},
	{"radek", "ramek", 1},
	{"", "", 0},
	{"same", "same", 0},
	{"asdfdsf", "radeksdfdssdasd", 9},
	{"radeksdfdssdasd", "asdfdsf", 9},
	{"was", "wadassdfs3qrffsdfdvb", 17},
}

func TestLevenshtein(t *testing.T) {

	for _, tc := range testCases {
		dist := strdist.Levenshtein(tc.source, tc.target)

		if dist != tc.distance {
			t.Errorf("In(%s, %s) = %d; want %d", tc.source, tc.target, dist, tc.distance)
		}
	}
}

var result int

func benchmarkLevenshtein(source, target string, b *testing.B) {
	var dist int
	for i := 0; i < b.N; i++ {
		dist = strdist.Levenshtein(source, target)
	}
	result = dist
}

func BenchmarkLevenshtein1(b *testing.B)  { benchmarkLevenshtein("a", "b", b) }
func BenchmarkLevenshtein2(b *testing.B)  { benchmarkLevenshtein("aa", "bb", b) }
func BenchmarkLevenshtein3(b *testing.B)  { benchmarkLevenshtein("aaa", "bbb", b) }
func BenchmarkLevenshtein10(b *testing.B) { benchmarkLevenshtein("aaaaaaaaaa", "bbbbbbbbbb", b) }
func BenchmarkLevenshtein20(b *testing.B) {
	benchmarkLevenshtein("aaaaaaaaaaaaaaaaaaaa", "bbbbbbbbbbbbbbbbbbbb", b)
}
func BenchmarkLevenshtein40(b *testing.B) {
	benchmarkLevenshtein("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", b)
}
