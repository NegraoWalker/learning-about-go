package enderecos

import "testing"

// Testes Unitários:
func TestTipoDeEndereco(t *testing.T) {
	enderecoParaTeste := "Avenida Maringá"
	tipoDeEnderecoEsperado := "Avenida"
	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
		t.Errorf("O tipo de endereço recebido é diferente do esperado. Esperava %s e foi recebido %s",
			tipoDeEnderecoEsperado, tipoDeEnderecoRecebido)
	}
}
