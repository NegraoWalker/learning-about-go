/*

Crie uma variável inteira x e um ponteiro p que aponte para x. Atualize o valor de x através de p.

*/

package main

import "fmt"

func main() {
	var x int = 100
	var p *int = &x

	fmt.Println("Endereço de memória da variável x: ", &x)
	fmt.Println("Valor armazenado no endereço de memória da variável x: ", x)
	fmt.Println("Endereço de memória apontado pela variável p: ", p)

	*p = 10 //Alterando o valor armazenado com a notação * da variável x pelo ponteiro
	fmt.Println("Valor armazenado no endereço de memória da variável x, após a alteração: ", x)
}
