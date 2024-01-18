package exercises

import "fmt"

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

func ExerciseThree() {

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

	fmt.Println(cam)
	fmt.Println(cam.veiculo)
	fmt.Println(sed)
	fmt.Println(sed.veiculo)
}
