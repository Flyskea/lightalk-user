package svc

import (
	"fmt"

	"github.com/flyskea/lightalk-user-rpc/internal/config"
	"github.com/flyskea/lightalk-user-rpc/model"
	"github.com/flyskea/lightalk-user-rpc/model/ent"
)

type ServiceContext struct {
	Config config.Config

	Ent *ent.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	dbInfo := c.DBInfo
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True",
		dbInfo.User, dbInfo.Password, dbInfo.IP, dbInfo.Port, dbInfo.DatabaseName)
	return &ServiceContext{
		Config: c,
		Ent:    model.InitDB(dsn),
	}
}
