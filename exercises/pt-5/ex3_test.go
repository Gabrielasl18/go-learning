package exercises

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestExerciseThree(t *testing.T) {
	var buf bytes.Buffer
	expected := `{{4 preto} false}
{4 preto}
{{2 branco} true}
{2 branco}` + "\n"

	exerciseThree(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

type veiculo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veiculo
	traçãoNasQuatro bool
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func exerciseThree(out io.Writer) {

	cam := caminhonete{
		veiculo: veiculo{
			portas: 4,
			cor:    "preto",
		},
		traçãoNasQuatro: false,
	}
	sed := sedan{
		veiculo: veiculo{
			portas: 2,
			cor:    "branco",
		},
		modeloLuxo: true,
	}

	fmt.Fprintln(out, cam)
	fmt.Fprintln(out, cam.veiculo)
	fmt.Fprintln(out, sed)
	fmt.Fprintln(out, sed.veiculo)
}
