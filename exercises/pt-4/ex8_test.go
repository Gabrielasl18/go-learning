package exercises

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestExerciseEight(t *testing.T) {
	var buf bytes.Buffer
	expected := "gabriela_lacerda \n\t0 - skate \n\t1 - ser legal" + "\n" + "luciana_silva \n\t0 - tocar viola \n\t1 - ser legal" +
		"\n" + "joana_pinheiro \n\t0 - surfar \n\t1 - ser chata\n"

	exerciseEight(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}

}

func exerciseEight(out io.Writer) {
	m := map[string][]string{
		"gabriela_lacerda": []string{
			"skate", "ser legal",
		},
		"luciana_silva": []string{
			"tocar viola", "ser legal",
		},
		"joana_pinheiro": []string{
			"surfar", "ser chata",
		},
	}

	for key, value := range m {
		fmt.Fprintln(out, key)
		for i, h := range value {
			fmt.Fprintln(out, "\t", i, "-", h)
		}
	}
}
