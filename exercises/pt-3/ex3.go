package exercises

import "fmt"

func ExerciseThree() {
	anonasc := 2003
	anoatual := 2023
	for anonasc <= anoatual {
		fmt.Println(anonasc)
		anonasc++
	}
}
