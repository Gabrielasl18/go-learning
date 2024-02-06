package exercises

import (
	"fmt"
	"sort"
)

type user_p struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ordenaPorIdade []user_p
type ordenaPorSobrenome []user_p

func (u ordenaPorIdade) Len() int {
	return len(u)
}
func (u ordenaPorIdade) Less(i, j int) bool {
	return u[i].Age < u[j].Age
}
func (u ordenaPorIdade) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}
func (u ordenaPorSobrenome) Len() int {
	return len(u)
}
func (u ordenaPorSobrenome) Less(i, j int) bool {
	return u[i].Last < u[j].Last
}
func (u ordenaPorSobrenome) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func ExerciseFive() {
	u1 := user_p{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user_p{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user_p{
		First: "M",
		Last:  "Asss",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user_p{u1, u2, u3}

	sort.Sort(ordenaPorSobrenome(users))
	fmt.Println(users)
}
