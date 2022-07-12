package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var Conn *sql.DB

// Driver名
const driverName = "postgres"

func init() {
	// ユーザ
	user := os.Getenv("POSTGRE_USER")
	// パスワード
	password := os.Getenv("POSTGRE_PASSWORD")
	// 接続先ホスト
	host := os.Getenv("POSTGRE_HOST")
	// 接続先ポート
	port := os.Getenv("POSTGRE_PORT")
	// 接続先データベース
	database := os.Getenv("POSTGRE_DATABASE")

	var err error
	Conn, err = sql.Open(driverName,
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database))
	if err != nil {
		log.Fatal(err)
	}

}
