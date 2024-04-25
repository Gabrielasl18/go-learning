package interface_go

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"testing"
)

func TestInterface_go(t *testing.T) {
	var buf bytes.Buffer
	expected := `{3 4}
12
14
{5}
78.53981633974483
31.41592653589793` + "\n"

	interface_go(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry, out io.Writer) {
	fmt.Fprintln(out, g)
	fmt.Fprintln(out, g.area())
	fmt.Fprintln(out, g.perim())
}

func interface_go(out io.Writer) {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r, out)
	measure(c, out)
}
