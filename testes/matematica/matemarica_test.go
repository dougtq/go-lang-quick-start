package matematica

import "testing"

const errPadrao = "Valor esperado %v, mas o ersultado encontrado foi %v."

func TestMedia(t *testing.T) {
	valorEsperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 7.9)

	if valor != valorEsperado {
		t.Errorf(errPadrao, valorEsperado, valor)
	}
}
