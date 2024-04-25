package exercises

import (
	"bytes"
	"fmt"
	"io"
	"sort"
	"testing"
)

func TestExerciseFour(t *testing.T) {
	var buf bytes.Buffer
	expected := `[5 8 2 43 17 987 14 12 21 1 4 2 3 93 13]
[1 2 2 3 4 5 8 12 13 14 17 21 43 93 987]
[random rainbow delights in torpedo summers under gallantry fragmented moons across magenta]
[across delights fragmented gallantry in magenta moons rainbow random summers torpedo under]` + "\n"

	exerciseFour(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

func exerciseFour(out io.Writer) {

	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Fprintln(out, xi)
	sort.Ints(xi)
	fmt.Fprintln(out, xi)

	fmt.Fprintln(out, xs)
	sort.Strings(xs)
	fmt.Fprintln(out, xs)
}
