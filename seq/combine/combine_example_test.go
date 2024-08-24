package combine_test

import (
	"errors"
	"fmt"
	"maps"
	"slices"
	"strconv"

	"github.com/GodsBoss/g/seq/combine"
	"github.com/GodsBoss/g/seq/iterate"
	"github.com/GodsBoss/g/seq/iterate2"
)

// ExampleTransformers_evenNumbersAsString combines a func(iter.Seq) iter.Seq and a func(iter.Seq) iter.Seq.
func ExampleTransformers_evenNumbersAsString() {
	isEven := func(n int) bool {
		return n%2 == 0
	}

	evenToString := combine.Transformers(iterate.Filter(isEven), iterate.Map(strconv.Itoa))

	for s := range evenToString(slices.Values([]int{1, 2, 3, 4, 5, 6})) {
		fmt.Printf("%s\n", s)
	}

	// Output:
	// 2
	// 4
	// 6
}

// ExampleTransformers_fetchResponses combines a func(iter.Seq2) iter.Seq2 and a func(iter.Seq2) iter.Seq2.
func ExampleTransformers_fetchResponses() {
	urls := maps.All(
		map[string]error{
			"https://foo.example.com": errors.New("unreachable"),
			"https://bar.example.com": nil,
			"https://baz.example.com": nil,
		},
	)

	filterSuccessfulCalls := iterate2.Filter(
		func(_ string, err error) bool {
			return err == nil
		},
	)

	fetchResponse := iterate2.Map2(
		func(url string, _ error) (contents string, length int) {
			contents = fmt.Sprintf("contents of %s", url)
			length = len(contents)
			return
		},
	)

	for contents, length := range combine.Transformers(filterSuccessfulCalls, fetchResponse)(urls) {
		fmt.Printf("%s (%d)\n", contents, length)
	}

	// Unordered output:
	// contents of https://bar.example.com (35)
	// contents of https://baz.example.com (35)
}

// ExampleTransformers_showNameLengthInsteadOfAge combines a func(iter.Seq2) iter.Seq and a func(iter.Seq) iter.Seq2.
func ExampleTransformers_showNameLengthInsteadOfAge() {
	people := maps.All(
		map[string]int{
			"Peter": 37,
			"Paul":  22,
			"Mary":  31,
		},
	)

	ignoreAge := iterate.OverFirst[string, int]

	addNameLength := iterate2.Map(
		func(name string) (string, int) {
			return name, len(name)
		},
	)

	for name, length := range combine.Transformers(ignoreAge, addNameLength)(people) {
		fmt.Printf("%s (%d)\n", name, length)
	}

	// Unordered output:
	// Peter (5)
	// Paul (4)
	// Mary (4)
}

// ExampleTransformers_sumQuotientAndRemainder combines a func(iter.Seq) iter.Seq2 and a func(iter.Seq2) iter.Seq.
func ExampleTransformers_sumQuotientAndRemainder() {
	divideByThree := iterate2.Map(
		func(number int) (quotient int, remainder int) {
			return number / 3, number % 3
		},
	)

	sumPair := iterate.Map2(
		func(n int, m int) int {
			return n + m
		},
	)

	for sum := range combine.Transformers(divideByThree, sumPair)(slices.Values([]int{2, 4, 7, 9, 11})) {
		fmt.Println(sum)
	}

	// Output:
	// 2
	// 2
	// 3
	// 3
	// 5
}
