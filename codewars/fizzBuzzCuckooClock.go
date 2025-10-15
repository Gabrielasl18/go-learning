package codewars

import (
	"strconv"
	"strings"
	"time"
)

/*

Quando os minutos são divisíveis por três, o relógio deve dizer "Fizz".
Quando os minutos são divisíveis por cinco, o relógio deve dizer "Buzz".
Se os minutos forem divisíveis por ambos três e cinco, o relógio deve dizer "Fizz Buzz", exceto em duas situações:
- Quando for hora cheia, a porta do relógio se abrirá, e o cuco sairá e dirá "Cuckoo" entre uma e doze vezes,
dependendo da hora.
- Quando for meia hora, a porta do relógio se abrirá, e o cuco sairá e dirá "Cuckoo" apenas uma vez.
- Se os minutos não forem divisíveis por três ou cinco, o relógio emitirá um som de "tick" para representar o
tempo passando de forma mais sutil.
- A entrada para a função será uma string contendo a hora e os minutos no formato de 24 horas, separados por
 dois pontos e com zeros à esquerda, como "13:34" para 1:34 da tarde.

O valor de retorno da função será uma string contendo a combinação dos sons "Fizz", "Buzz", "Cuckoo" e/ou "tick"
que o relógio precisa fazer naquele momento, separados por espaços. Note que, embora a entrada esteja em formato
de 24 horas, os "Cuckoos" do relógio cuco seguem o formato de 12 horas.

*/

/*

se os min divisiveis por 3 == fizz
se os min divisiveis por 5 == buzz
se os min divisiveis por 3 e 5 == fizz buzz
se hora cheia (12:00/02:00) == cuckoo, entre 1 e 12 vezes, dependendo da hora
se for meia hora (12:30/02:30) == cuckoo, só uma vez
se nao for divisivel por 3 e 5, saíra tick
*/

// input -> codewars.FizzBuzzCuckooClock("22:00")
func FizzBuzzCuckooClock(timeString string) string {
	clockString := strings.Split(timeString, ":")
	hourInt, _ := strconv.Atoi(clockString[0])
	minuteInt, _ := strconv.Atoi(clockString[1])
	var word string

	time24 := time.Date(2000, 1, 1, hourInt, minuteInt, 0, 0, time.UTC)

	hour12 := time24.Format("3")
	time12, _ := strconv.Atoi(hour12)

	switch {
	case minuteInt != 0 && minuteInt != 30:
		if minuteInt%3 == 0 && minuteInt%5 == 0 {
			word = "Fizz Buzz"
		} else if minuteInt%3 == 0 {
			word = "Fizz"
		} else if minuteInt%5 == 0 {
			word = "Buzz"
		} else {
			word = "tick"
		}
	case minuteInt == 0:
		if hourInt != 0 && hourInt != 12 {
			for i := 0; i < time12; i++ {
				if i+1 != time12 {
					word += "Cuckoo "
				} else {
					word += "Cuckoo"
				}
			}
		} else {
			for i := 0; i < 12; i++ {
				if i+1 != 12 {
					word += "Cuckoo "
				} else {
					word += "Cuckoo"
				}
			}
		}
	case minuteInt == 30:
		word = "Cuckoo"
	default:
		word = "tick"
	}
	return word
}
