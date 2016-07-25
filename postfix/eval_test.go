package postfix

import (
	"testing"
)

func TestEvaluate(t *testing.T) {

	data := []struct {
		exp    []rune
		evar   map[rune]bool
		eval   []rune
		result bool
	}{
		{[]rune{'a', '!', 'a', 'a', '&', '|'}, map[rune]bool{'a': true}, []rune{'1'}, true},
		{[]rune{'a', '!', 'b', 'a', '!', '&', '|'}, map[rune]bool{'a': true, 'b': true}, []rune{'1', '1'}, false},
		{[]rune{'a', '!', 'a', '|'}, map[rune]bool{'a': true}, []rune{'0'}, true},
		{[]rune{'a', 'b', '!', 'b', '|', '&', 'a', '!', 'b', '!', 'b', '|', '&', '|'},
			map[rune]bool{'a': true, 'b': true}, []rune{'0', '1'}, true},
		{[]rune{'a', 'b', '!', '|'}, map[rune]bool{'a': true, 'b': true}, []rune{'1', '1'}, true},
	}

	for _, e := range data {
		res := Evaluate(e.exp, e.evar, e.eval)
		if res != e.result {
			t.Errorf("Got: %v, Expected: %v\n", res, e.result)
		}
	}

}
