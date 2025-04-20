package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Server struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"server"`
	Mysql struct {
		Host string `json:"host"`
		User string `json:"user"`
		Pass string `json:"pass"`
		Name string `json:"name"`
	} `json:"mysql"`
	Discord struct {
		Webhook   string `json:"webhook"`
		Image     string `json:"image"`
		Website   string `json:"website"`
	} `json:"discord"`
	CncServer struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"cnc_server"`
	SSHMethods []struct {
		Name    string `json:"name"`
		Command string `json:"command"`
	} `json:"ssh_methods"`
	MiraiMethods []struct {
		Name    string `json:"name"`
		Command string `json:"command"`
	} `json:"mirai_methods"`
	QbotMethods []struct {
		Name    string `json:"name"`
		Command string `json:"command"`
	} `json:"qbot_methods"`
	LicenseKey string `json:"license"`
}

var Cfg *Config

func LoadConfig(path string) error {
	log.Printf("\u001B[0m\u001B[107m\u001B[38;5;163m[load/config]\u001B[0m\u001B[38;5;046m Loading! \u001B[38;5;230m")
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &Cfg)
	if err != nil {
		return err
	}
	log.Printf("\u001B[0m\u001B[107m\u001B[38;5;163m[load/config]\u001B[0m\u001B[38;5;046m(config) Loaded! \u001B[38;5;230m")
	return nil
}
