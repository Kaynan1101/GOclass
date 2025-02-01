package formas

import (
	"math"
	"testing"
)

// aqui mostra o teste de uma função em diferentes tipos de calculos e o T.Run
func TestArea(t *testing.T) {
	t.Run("Retângulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			//fatal = o codigo não prossegue, erro prossegue
			t.Fatalf("Área recebida: %f é diferente da área esperada: %f", areaRecebida, areaEsperada)
		}

	})

	t.Run("Círculo", func(t *testing.T) {
		circ := Circulo{10}

		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circ.Area()

		if areaEsperada != areaRecebida {
			//fatal = o codigo não prossegue, erro prossegue
			t.Fatalf("Área recebida: %f é diferente da área esperada: %f", areaRecebida, areaEsperada)
		}

	})
}
