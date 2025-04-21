package auxiliar

import "fmt"

func EscreverArquivo1() { //Nome da função começando com letra maiúsculo - Pode ser exportada e usada em outros pacotes
	fmt.Println("arquivo1.go")
	escreverArquivo2()
}
