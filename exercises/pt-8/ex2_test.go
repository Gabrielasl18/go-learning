package exercises

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"testing"
)

func TestExerciseTwo(t *testing.T) {
	var buf bytes.Buffer
	expected := `<nil>
[{James Bond 32 [Shaken, not stirred Youth is no guarantee of innovation In his majesty's royal service]} {Miss Moneypenny 27 [James, it is soo good to see you Would you like me to take care of that for you, James? I would really prefer to be a secret agent myself.]} {M Hmmmm 54 [Oh, James. You didn't. Dear God, what has James done now? Can someone please tell me where James Bond is?]}]` + "\n"

	exerciseTwo(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

type User2 []struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func exerciseTwo(out io.Writer) {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	var userTest User2
	b := json.Unmarshal([]byte(s), &userTest)
	fmt.Fprintln(out, b)
	fmt.Fprintln(out, userTest)

}
