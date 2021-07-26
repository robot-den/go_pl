package word

import "testing"

func TestPalindrome(t *testing.T) {
	if !IsPalindrome("detartrated") {
		t.Error(`IsPalindrome("detartrated") == false`)
	}

	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrome("kayak") == false`)
	}

	if !IsPalindrome("été") {
		t.Error(`IsPalindrome("été") == false`)
	}
}

func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("word") {
		t.Error(`IsPalindrome("word") == true`)
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N ; i++ {
		IsPalindrome("Oh my dear, please go sleep, I'm beggin you!")
	}
}
