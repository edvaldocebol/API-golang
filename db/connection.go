package db

import (
	"database/sql"
	"fmt"

	"github.com/edvaldocebol/API-golang.git/configs"
	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()
	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	conn, err := sql.Open("postgres", sc)
	if err != nil {
		return nil, fmt.Errorf("Erro ao abrir a conexão com o banco de dados: %v", err)
	}

	if err = conn.Ping(); err != nil {
		conn.Close()
		return nil, fmt.Errorf("Erro ao verificar a conexão com o banco de dados: %v", err)
	}

	return conn, nil
}
