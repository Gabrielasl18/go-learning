package exercises

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestExerciseFive(t *testing.T) {
	var buf bytes.Buffer
	expected := "gabriela lacerda\n"
	exerciseFour(&buf)
	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

type figura interface {
	area() float64
}

type quadrado struct {
	lado float64
}

type circulo struct {
	raio float64
}

func (q quadrado) area() float64 {
	return q.lado * q.lado
}

func (c circulo) area() float64 {
	return 2 * 3.14 * c.raio
}

func info(f figura, out io.Writer) {
	fmt.Fprintln(out, f.area())
}

func ExerciseFive(out io.Writer) {
	Circulo := circulo{
		raio: 2,
	}
	Quadrado := quadrado{
		lado: 3,
	}
	info(Circulo, out)
	info(Quadrado, out)
}
