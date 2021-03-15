package main

import "testing"

// TestBalanced вызывает isBracketsBalanced со сбалансированными строками
func TestBalanced(t *testing.T) {
	strings := []string{
		"",
		"s",
		"()",
		"[[a](b){((c))}[d]](e)",
	}
	for _, str := range strings {
		if balanced, pos := isBracketsBalanced(str); !balanced {
			t.Fatalf(`isBalanced("%s") = %t, %d, want true`, str, balanced, pos)
		}
	}
}

// TestUnBalanced вызывает isBracketsBalanced с несбалансированными строками
func TestUnBalanced(t *testing.T) {
	strings := []string{
		"(p",
		"]",
		"d)",
		"[[a](b)){((c))}[d]](e)",
	}
	for _, str := range strings {
		if balanced, pos := isBracketsBalanced(str); balanced {
			t.Fatalf(`isBalanced("%s") = %t, %d, want false`, str, balanced, pos)
		}
	}
}
