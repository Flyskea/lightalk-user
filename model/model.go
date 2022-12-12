package model

import "github.com/flyskea/lightalk-user-rpc/model/ent"

func InitDB(dsn string) *ent.Client {
	client, err := ent.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return client
}
