package exercises

import "fmt"

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

func info(f figura) {
	fmt.Println(f.area())
}

func ExerciseFive() {
	Circulo := circulo{
		raio: 2,
	}
	Quadrado := quadrado{
		lado: 3,
	}
	info(Circulo)
	info(Quadrado)
}
