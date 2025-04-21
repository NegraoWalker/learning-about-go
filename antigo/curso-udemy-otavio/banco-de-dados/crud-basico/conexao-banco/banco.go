package conexaobanco

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //Driver de conexão com o MySQL
)

func RealizaConexaoComOBanco() (*sql.DB, error) {
	stringDeConexao := "root:12345@tcp(127.0.0.1:3306)/devbook?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", stringDeConexao)

	if erro != nil {
		fmt.Println("Erro gerado pelo sql.Open(): ")
		return nil, erro
	}

	erro = db.Ping() //Testando se a conexão com o banco deu certo

	if erro != nil {
		fmt.Println("Erro gerado pelo db.Ping(): ")
		return nil, erro
	}

	fmt.Println("Conexão com o banco efetuada com sucesso!!")
	return db, nil
}
