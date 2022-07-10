package model

import (
	"OcSummerAPI/pkg/database"
	"OcSummerAPI/pkg/domain"
	"crypto/sha256"
	"fmt"
)

// sidechannel

type Login struct {
	ID       string
	Name     string
	PassWord string
}

// 駄目なログインフォームのバックエンドコード
func SearchUser(request *domain.LoginRequest) error {
	login := Login{}
	row := database.Conn.QueryRow("SELECT * FROM user WHERE name = ?", login.Name)
	if err := row.Scan(&login.ID, &login.Name, &login.PassWord); err != nil {
		return fmt.Errorf("ユーザーが存在しません")
	}
	sum := fmt.Sprintf("%x", sha256.Sum256([]byte(request.Pass)))
	if login.PassWord != sum {
		return fmt.Errorf("パスワードが違います")
	}
	return nil
}
