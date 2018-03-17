package strings

import "testing"
import "strings"

const msgIndex = "%s (parte: %s) - Indices: esperado (%d) <> encontrado (%d)."

func TestIndex(t *testing.T) {
	testes := []struct {
		texto    string
		parte    string
		esperado int
	}{
		{"Beblue é Show", "Beblue", 0},
		{"", "", 0},
		{"Opa é Show", "opa", -1},
		{"Itor é Show", "t", 1},
	}

	for _, teste := range testes {
		t.Logf("Massa: %v", teste)
		atual := strings.Index(teste.texto, teste.parte)

		if atual != teste.esperado {
			t.Errorf(msgIndex, teste.texto, teste.parte, teste.esperado, atual)
		}
	}
}
