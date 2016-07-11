package postfix

import (
	"testing"
)

type expression struct {
	infix, postfix string
}

func TestConvert(t *testing.T) {

	exp := []expression{
		{"(!a|(a&a))", "a!aa&|"},
		{"(!a|(b&!a))", "a!ba!&|"},
		{"(!a|a)", "a!a|"},
		{"((a&(!b|b))|(!a&(!b|b)))", "ab!b|&a!b!b|&|"},
	}

	for _, e := range exp {
		p, _ := Convert(e.infix)
		if string(p) != e.postfix {
			t.Errorf("Got: %s, Expected: %s\n", string(p), e.postfix)
		}
	}

}
