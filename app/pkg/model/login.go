package model

import (
	"OCsemmerApp/pkg/db"
	"OCsemmerApp/pkg/domain"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"log"
)

// sidechannel

type Login struct {
	ID       string
	Name     string
	PassWord string
}

// 駄目なログインフォームのバックエンドコード
func SearchUser(request domain.LoginRequest) error {
	row := db.Conn.QueryRow("SELECT * FROM users WHERE name = $1", request.Name)
	login := Login{}
	// ユーザーがいない場合パスワード検証をせず返してしまう
	if err := row.Scan(&login.ID, &login.Name, &login.PassWord); err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("ユーザーが存在しません")
		}
		log.Println(err)
		return nil
	}
	sum := fmt.Sprintf("%x", sha256.Sum256([]byte(request.Pass)))
	if login.PassWord != sum {
		return fmt.Errorf("パスワードが違います")
	}
	return nil
}
