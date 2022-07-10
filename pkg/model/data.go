package model

import (
	"OcSummerAPI/pkg/database"
	"database/sql"
	"fmt"
)

type Data struct {
	ProductID    string
	ProductName  string
	UserName     string
	PurchaseDate string
}

type DataList []*Data

// ユーザーが入力した名前で検索できてしまうやつ
func SearchData(name string) (*DataList, error) {
	rows, err := database.Conn.Query(
		fmt.Sprintf(
			`SELECT
				T1.product_id
				T1.product_name
				T2.name
				T1.purchase_date
			FROM 
				purchase AS T1
				JOIN user AS T2
					ON T1.user_id = T2.id
			WHERE
				T2.name = %s`, name))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var datalist DataList
	for rows.Next() {
		data, err := convertToData(rows)
		if err != nil {
			return nil, err
		}
		datalist = append(datalist, data)
	}
	return &datalist, nil
}

func convertToData(rows *sql.Rows) (*Data, error) {
	data := Data{}
	if err := rows.Scan(
		&data.ProductID,
		&data.ProductName,
		&data.UserName,
		&data.PurchaseDate); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &data, nil
}
