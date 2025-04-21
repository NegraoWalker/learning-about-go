package servidor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	conexaobanco "mod-crud-basico/conexao-banco"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func BuscaUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := conexaobanco.RealizaConexaoComOBanco()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}

	defer db.Close()

	linhas, erro := db.Query("SELECT * FROM usuarios")
	if erro != nil {
		w.Write([]byte("Erro ao buscar os usuários!"))
		return
	}

	defer linhas.Close()

	var usuarios []usuario
	for linhas.Next() {
		var usuario usuario

		erro := linhas.Scan(&usuario.Id, &usuario.Nome, &usuario.Email)
		if erro != nil {
			w.Write([]byte("Erro ao escanear o usuário!"))
			return
		}
		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)
	erro = json.NewEncoder(w).Encode(usuarios)
	if erro != nil {
		w.Write([]byte("Erro ao converter os usuários para JSON!"))
		return
	}
}

func BuscaUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	Id, erro := strconv.ParseInt(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro!"))
		return
	}

	db, erro := conexaobanco.RealizaConexaoComOBanco()

	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}

	linha, erro := db.Query("SELECT * FROM usuarios WHERE id = ?", Id)

	if erro != nil {
		w.Write([]byte("Erro ao buscar o usuário!"))
		return
	}

	var usuario usuario
	if linha.Next() {
		erro := linha.Scan(&usuario.Id, &usuario.Nome, &usuario.Email)
		if erro != nil {
			w.Write([]byte("Erro ao escanear o usuário!"))
			return
		}
	}

	w.WriteHeader(http.StatusOK)

	erro = json.NewEncoder(w).Encode(usuario)
	if erro != nil {
		w.Write([]byte("Erro ao converter os usuários para JSON!"))
		return
	}
}

func AtualizaUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	Id, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		http.Error(w, "Erro ao converter o parâmetro para inteiro!", http.StatusBadRequest)
		return
	}

	corpoDaRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		http.Error(w, "Erro ao ler corpo da requisição!", http.StatusBadRequest)
		return
	}

	var usuario usuario
	if erro := json.Unmarshal(corpoDaRequisicao, &usuario); erro != nil {
		http.Error(w, "Erro ao converter o usuário para struct!", http.StatusBadRequest)
		return
	}

	db, erro := conexaobanco.RealizaConexaoComOBanco()
	if erro != nil {
		http.Error(w, "Erro ao conectar no banco de dados!", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("UPDATE usuarios SET nome = ?, email = ? WHERE id = ?")
	if erro != nil {
		http.Error(w, "Erro ao criar o statement!", http.StatusInternalServerError)
		return
	}
	defer statement.Close()

	_, erro = statement.Exec(usuario.Nome, usuario.Email, Id)
	if erro != nil {
		http.Error(w, "Erro ao atualizar o usuário!", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeletaUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	Id, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		http.Error(w, "Erro ao converter o parâmetro para inteiro!", http.StatusBadRequest)
		return
	}
	db, erro := conexaobanco.RealizaConexaoComOBanco()
	if erro != nil {
		http.Error(w, "Erro ao conectar no banco de dados!", http.StatusInternalServerError)
		return
	}
	defer db.Close()
	statement, erro := db.Prepare("DELETE FROM usuarios WHERE id = ?")
	if erro != nil {
		http.Error(w, "Erro ao criar o statement!", http.StatusInternalServerError)
		return
	}
	defer statement.Close()
	_, erro = statement.Exec(Id)
	if erro != nil {
		http.Error(w, "Erro ao deletar o usuário!", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
