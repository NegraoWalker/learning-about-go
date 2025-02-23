package servidor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	conexaobanco "mod-crud-basico/conexao-banco"
	"net/http"
)

type usuario struct {
	Id    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriaUsuario(w http.ResponseWriter, r *http.Request) {
	corpoDaRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}
	var usuario usuario
	erro = json.Unmarshal(corpoDaRequisicao, &usuario)
	if erro != nil {
		w.Write([]byte("Erro em converter o usuario para struct!"))
		return
	}
	fmt.Println(usuario)

	db, erro := conexaobanco.RealizaConexaoComOBanco()
	if erro != nil {
		w.Write([]byte("Erro ao tentar conexão com o banco de dados!"))
		return
	}
	statement, erro := db.Prepare("INSERT INTO usuarios (nome, email) VALUES (?,?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}

	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement"))
		return
	}

	idCriado, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o Id criado!"))
		return
	}

	w.WriteHeader(http.StatusCreated) //201
	w.Write([]byte(fmt.Sprintf("usuario inserido com sucesso! Id: %d", idCriado)))
}
