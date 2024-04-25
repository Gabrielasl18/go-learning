package loop

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestLooping(t *testing.T) {
	var buf bytes.Buffer
	expected := `0 x 0 = 0
1 x 0 = 0
2 x 0 = 0
3 x 0 = 0
4 x 0 = 0
5 x 0 = 0
6 x 0 = 0
7 x 0 = 0
8 x 0 = 0
9 x 0 = 0
0 x 1 = 0
1 x 1 = 1
2 x 1 = 2
3 x 1 = 3
4 x 1 = 4
5 x 1 = 5
6 x 1 = 6
7 x 1 = 7
8 x 1 = 8
9 x 1 = 9
0 x 2 = 0
1 x 2 = 2
2 x 2 = 4
3 x 2 = 6
4 x 2 = 8
5 x 2 = 10
6 x 2 = 12
7 x 2 = 14	
8 x 2 = 16
9 x 2 = 18` + "\n"

	looping(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

func looping(out io.Writer) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 10; j++ {
			fmt.Fprintf(out, "%d x %d = %d\n", j, i, j*i)
		}
	}
}
