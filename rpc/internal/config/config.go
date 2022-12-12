package config

import "github.com/zeromicro/go-zero/zrpc"

type DBInfo struct {
	IP           string `env:"IP"`
	Port         int    `env:"PORT" envDefault:"3306"`
	User         string `env:"USER" envDefault:"root"`
	Password     string `env:"PASSWORD"`
	DatabaseName string `env:"DATABASENAME"`
}

type Config struct {
	zrpc.RpcServerConf

	DBInfo *DBInfo
}
