package model

import (
	"github.com/flyskea/lightalk-user/model/ent"
	_ "github.com/go-sql-driver/mysql"
)

func InitDB(dsn string) *ent.Client {
	client, err := ent.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return client
}
