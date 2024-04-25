package exercises

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestExerciseThree(t *testing.T) {
	var buf bytes.Buffer
	expected := `[{James Bond 32 [Shaken, not stirred Youth is no guarantee of innovation In his majesty's royal service]} {Miss Moneypenny 27 [James, it is soo good to see you Would you like me to take care of that for you, James? I would really prefer to be a secret agent myself.]} {M Hmmmm 54 [Oh, James. You didn't. Dear God, what has James done now? Can someone please tell me where James Bond is?]}]
0` + "\n"

	exerciseThree(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

type user3 struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func exerciseThree(out io.Writer) {
	u1 := user3{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user3{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user3{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user3{u1, u2, u3}

	fmt.Fprintln(out, users)

	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintln(out, "0")

}
