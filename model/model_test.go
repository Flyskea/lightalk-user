package model_test

import (
	"context"
	"log"
	"testing"

	"github.com/flyskea/lightalk-user-rpc/model/ent"
	_ "github.com/go-sql-driver/mysql"
)

func TestModel(t *testing.T) {
	client, err := ent.Open("mysql", "root:Flyskea.0830@tcp(www.1696666.xyz:3306)/forum?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// 运行自动迁移工具。
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
