package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	input  Input
	output bool
}

type Input struct {
	a []any
	b []any
}

func TestInOrder(t *testing.T) {
	testCases := []TestCase{
		{input: Input{a: []any{float64(1), float64(1), float64(3), float64(1), float64(1)}, b: []any{float64(1), float64(1), float64(5), float64(1), float64(1)}}, output: true},

		{input: Input{a: []any{[]any{float64(1)}, []any{float64(2), float64(3), float64(4)}}, b: []any{[]any{float64(1)}, float64(4)}}, output: true},

		{input: Input{a: []any{float64(9)}, b: []any{[]any{float64(8), float64(7), float64(6)}}}, output: false},

		{input: Input{a: []any{[]any{float64(4), float64(4)}, float64(4), float64(4)}, b: []any{[]any{float64(4), float64(4)}, float64(4), float64(4), float64(4)}}, output: true},

		{input: Input{a: []any{float64(7), float64(7), float64(7), float64(7)}, b: []any{float64(7), float64(7), float64(7)}}, output: false},

		{input: Input{a: []any{}, b: []any{float64(3)}}, output: true},

		{input: Input{a: []any{[]any{[]any{}}}, b: []any{[]any{}}}, output: false},

		{input: Input{a: []any{float64(1), []any{float64(2), []any{float64(3), []any{float64(4), []any{float64(5), float64(6), float64(7)}}}}, float64(8), float64(9)}, b: []any{float64(1), []any{float64(2), []any{float64(3), []any{float64(4), []any{float64(5), float64(6), float64(0)}}}}}}, output: false},

		{input: Input{a: []any{[]any{float64(1), float64(8), float64(8)}, []any{float64(9)}}, b: []any{float64(9)}}, output: true},
	}

	for _, tc := range testCases {
		fmt.Println("NEW TEST")
		got := AreEqual(tc.input.a, tc.input.b)
		assert.Equal(t, tc.output, got, tc.input)
	}

}
