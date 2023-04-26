package main

import (
	"fmt"

	"gopkg.in/go-ini/ini.v1"
	_ "gopkg.in/go-ini/ini.v1"
)

//ini

type ConfigList struct {
	Port      int
	DbName    string
	SQLDriver string
}

var Config ConfigList

func init() {
	cfg, _ := ini.Load("config.ini")

	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustInt(8080),
		DbName:    cfg.Section("db").Key("name").MustString("example.sql"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
	}
}

func main() {
	fmt.Println(Config)
}
