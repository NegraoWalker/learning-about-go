package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	stringDeConexao := "root:12345@tcp(127.0.0.1:3306)/devbook?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", stringDeConexao)

	if erro != nil {
		log.Fatal("Erro sql.Open = ", erro)
	}

	defer db.Close()

	erro = db.Ping() //Testando se a conexão com o banco deu certo

	if erro != nil {
		log.Fatal("Erro db.Ping = ", erro)
	}

	fmt.Println("Conexão com o banco efetuada com sucesso!!")

	linhas, erro := db.Query("SELECT * FROM usuarios") //Consulta no banco

	if erro != nil {
		log.Fatal("Erro db.Query [SELECT * FROM usuarios] = ", erro)
	}

	defer linhas.Close()

	fmt.Println("Usuários: ", linhas)
}
