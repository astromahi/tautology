package postfix

import (
	"testing"
)

func TestConvert(t *testing.T) {

	exp := []struct {
		infix, postfix string
	}{
		{"(!a|(a&a))", "a!aa&|"},
		{"(!a|(b&!a))", "a!ba!&|"},
		{"(!a|a)", "a!a|"},
		{"((a&(!b|b))|(!a&(!b|b)))", "ab!b|&a!b!b|&|"},
		{"a|!b", "ab!|"},
	}

	for _, e := range exp {
		p, _ := Convert(e.infix)
		if string(p) != e.postfix {
			t.Errorf("Got: %s, Expected: %s\n", string(p), e.postfix)
		}
	}

}
