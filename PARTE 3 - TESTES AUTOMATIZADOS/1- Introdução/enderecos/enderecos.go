package enderecos

import "strings"

// TipoDeEnderecos = Teste manual se um endereço é valido ou não
func TipoDeEnderecos(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "rodovia", "estrada"}

	enderecoEmLetraMiniscula := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetraMiniscula, " ")[0]

	enderecoTemUmTipoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return strings.Title(primeiraPalavraDoEndereco)
	}

	return "Tipo Inválido"
}
