package enderecos

import "testing"

func TestTipoDeEndero(t *testing.T) {
	endecoParaTeste := "Avenida Paulista"
	tipoDeEnderecoEsperado := "Avenida"
	tipoDeEnderecoRecebido := TipoDeEndereco(endecoParaTeste)

	if tipoDeEnderecoRecebido !== titipoDeEnderecoEsperado {
		t.Error("O tipo recebido e diferente do esperado!")
	}
}
