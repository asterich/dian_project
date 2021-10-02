package utils

import (
	"log"

	"gopkg.in/ini.v1"
)

var (

	// Server configs
	AppMode  string
	HttpPort string

	// Database configs
	Db     string
	DbName string
	DbPath string
)

func init() {
	var file, err = ini.Load("config/config.ini")
	if err != nil {
		log.Fatalln("Failed to load config.ini, err:", err.Error())
	}
	LoadServer(file)
	LoadDb(file)
}

func LoadServer(file *ini.File) {
	var serverSection = file.Section("server")
	AppMode = serverSection.Key("AppMode").String()
	HttpPort = serverSection.Key("HttpPort").String()

}

func LoadDb(file *ini.File) {
	var dbSection = file.Section("database")
	Db = dbSection.Key("Db").String()
	DbName = dbSection.Key("DbName").String()
	DbPath = dbSection.Key("DbPath").String()
}
