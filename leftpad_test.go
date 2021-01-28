package leftpad

import "testing"

func TestLeftpadSpace(t *testing.T) {
	want := []rune("   Hello, world.")
	if got := LeftpadSpace([]rune("Hello, world."), 3); compare(got, want) {
		t.Errorf("LeftpadSpace() was %q, not %q", got, want)
	}
}

func TestLeftpad(t *testing.T) {
	want := []rune("+++Hello, world.")
	if got := Leftpad([]rune("Hello, world."), 3, '+'); compare(got, want) {
		t.Errorf("Leftpad() was %q, not %q", got, want)
	}
}

func TestRJust(t *testing.T) {
	want := []rune("                                                                 Hello, world.")
	if got := RJust([]rune("Hello, world."), 78); compare(got, want) {
		t.Errorf("Leftpad() was %q, not %q", got, want)
	}
}

func compare(got []rune, want []rune) bool {
	for i, v := range got {
		if want[i] != v {
			return true
		}
	}

	return false
}

                                                                 