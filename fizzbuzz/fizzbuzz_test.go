package fizzbuzz_test

import (
	"hello/fizzbuzz"
	"testing"
)

func TestGiven1Wants1(t *testing.T) {
	given := 1
	wants := "1"

	get := fizzbuzz.Say(given)
	if wants != get {
		t.Errorf("given %d wants %q but get %q\n", given, wants, given)
	}
}

func TestGiven3WantsFizz(t *testing.T) {
	given := 3
	wants := "Fizz"

	get := fizzbuzz.Say(given)
	if wants != get {
		t.Errorf("given %d wants %q but get %q\n", given, wants, given)
	}
}

func TestGiven5WantsBuzz(t *testing.T) {
	given := 5
	wants := "Buzz"

	get := fizzbuzz.Say(given)
	if wants != get {
		t.Errorf("given %d wants %q but get %q\n", given, wants, given)
	}
}

func TestGiven15WantsFizzBuzz(t *testing.T) {
	given := 15
	wants := "FizzBuzz"

	get := fizzbuzz.Say(given)
	if wants != get {
		t.Errorf("given %d wants %q but get %q\n", given, wants, given)
	}
}