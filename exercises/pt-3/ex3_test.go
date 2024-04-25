package exercises

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestExerciseThree(t *testing.T) {
	var buf bytes.Buffer
	expected := `2003
2004
2005
2006
2007
2008
2009
2010
2011
2012
2013
2014
2015
2016
2017
2018
2019
2020
2021
2022
2023` + "\n"

	exerciseThree(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

func exerciseThree(out io.Writer) {
	anonasc := 2003
	anoatual := 2023
	for anonasc <= anoatual {
		fmt.Fprintln(out, anonasc)
		anonasc++
	}
}
