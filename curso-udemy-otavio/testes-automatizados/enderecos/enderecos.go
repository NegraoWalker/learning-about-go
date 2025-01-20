package enderecos

import "strings"

// TipoDeEndereco verifica se um endereço tem um tipo válido e o retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"Rua", "Avenida", "Estrada", "Rodovia"}
	enderecoComLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Title(strings.Split(enderecoComLetraMinuscula, " ")[0]) // Capitaliza a primeira palavra

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			return tipo
		}
	}

	return "Tipo Inválido!!"
}
