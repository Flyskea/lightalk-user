package svc

import (
	"fmt"

	"github.com/flyskea/lightalk-user/model"
	"github.com/flyskea/lightalk-user/model/ent"
	"github.com/flyskea/lightalk-user/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config

	Ent *ent.Client
}

func NewServiceContext(c config.Config, cfg config.ConfigEnv) *ServiceContext {
	dbInfo := cfg.DBInfo
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True",
		dbInfo.User, dbInfo.Password, dbInfo.IP, dbInfo.Port, dbInfo.DatabaseName)
	return &ServiceContext{
		Config: c,
		Ent:    model.InitDB(dsn),
	}
}
