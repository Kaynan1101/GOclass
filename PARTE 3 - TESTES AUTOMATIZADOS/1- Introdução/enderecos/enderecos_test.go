package enderecos

import "testing"

//Teste com varios cenarios
type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEnderecos(t *testing.T) {
	// TESTE DE UNIDADE = ELE VAI TESTAR UMA PEQUENA PARTE DO SEU CÓDIGO
	// enderecoParaTeste := "rua ABC"
	// tipoDeEnderecoEsperado := "Avenida"
	// tipoDeEnderecoRecebido := TipoDeEnderecos(enderecoParaTeste)

	// if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
	// 	t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s", tipoDeEnderecoEsperado, tipoDeEnderecoRecebido)
	// }

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"avenida ABC", "Avenida"},
		{"ESTRADA ABC", "estrada"}, //FAILED = PELO FATO QUE O RETORNO DEVE SER COM COMEÇO EM LETRA MAISCULA QUE NEM COLOCOU NA FUNÇÃO "TITLE"
		{"", "Tipo Inválido"},
		{"Ponte ABC", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TipoDeEnderecos(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido: %s é diferente do esperado: %s", retornoRecebido, cenario.retornoEsperado)
		}
	}
}
