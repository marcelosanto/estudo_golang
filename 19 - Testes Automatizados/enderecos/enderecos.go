package enderecos

import "strings"

// TipoDeEndereco verifica se um endereco tem um tipo valido e o retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecoEmLetrasMinusculas := strings.ToLower(endereco)
	pimeiraPalavraDoendereco := strings.Split(enderecoEmLetrasMinusculas, " ")[0]

	enderecoTemUmTipoValido := false
	for _, tipo := range tiposValidos {
		if tipo == pimeiraPalavraDoendereco {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return pimeiraPalavraDoendereco
	}

	return "Tipo Inválido"
}
