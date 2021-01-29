package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver de conexão com o Mysql
)

// Conectar abre a conexão com o banco de dandos
func Conectar() (*sql.DB, error) {
	stringConexao := "LMnRRDPCp3:EfDnEPTKcX@tcp(remotemysql.com:3306)/LMnRRDPCp3?charset=utf8&parseTime=True"

	db, err := sql.Open("mysql", stringConexao)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
