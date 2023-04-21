package config

import (
	"Users/xushao/Desktop/Apps/Gos/shigou_jo/Go/section18-Todo/utils"
	"log"

	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port      string
	SQLDriver string
	Dbname    string
	LogFile   string
}

var Config ConfigList

func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}

	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		Dbname:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
	}
}
