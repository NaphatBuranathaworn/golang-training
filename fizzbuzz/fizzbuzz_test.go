package fizzbuzz_test

import (
	"hello/fizzbuzz"
	"testing"
)

func TestGiven1Wants1(t *testing.T) {
	given := 4
	wants := "4"

	get := fizzbuzz.Say(given)
	if wants != get {
		t.Errorf("given %d wants %q but get %q\n", given, wants, given)
	}
}