package exercises

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"testing"
)

func TestExerciseOne(t *testing.T) {
	var buf bytes.Buffer
	expected := `[{James 32} {Moneypenny 27} {M 54}]
[{"First":"James","Age":32},{"First":"Moneypenny","Age":27},{"First":"M","Age":54}]` + "\n"

	exerciseOne(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

type user struct {
	First string
	Age   int
}

func exerciseOne(out io.Writer) {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Fprintln(out, users)

	test, err := json.Marshal(users)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintln(out, string(test))
}
