package main

import (
	"fmt"
	"mod-testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Avenida Maringá")
	fmt.Println(tipoEndereco)

	tipoEndereco = enderecos.TipoDeEndereco("Patrimônio")
	fmt.Println(tipoEndereco)
}
