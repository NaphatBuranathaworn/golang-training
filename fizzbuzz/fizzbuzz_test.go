package fizzbuzz_test

import (
	"hello/fizzbuzz"
	"testing"
)

func TestGiven1Wants1(t *testing.T) {
	given := 3
	wants := "Fizz"

	get := fizzbuzz.Say(given)
	if wants != get {
		t.Errorf("given %d wants %q but get %q\n", given, wants, given)
	}
}