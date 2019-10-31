package conf

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var (
	conf string
)

type HttpServer struct {
	Address       string        `json:"address"`
	CheckInterval time.Duration `json:"checkInterval"`
	Tolerance     int           `json:"tolerance"`
	PushInterval  time.Duration `json:"pushInterval"`
}

type configuration struct {
	// TODO 容错     tolerance
	// TODO 推送间隔 pushInterval
	// TODO 推送列表 []serverChanSCK
	Server        []HttpServer `json:"httpServer"`
	ServerChanSCK []string     `json:"serverChanSCK"`
	LogMode       bool         `json:"logMode"`
}

var Config configuration

func init() {
	flag.StringVar(&conf, "conf", "", "configuration file")
	flag.Parse()
	if conf == "" {
		log.Fatal("no configuration file")
	}
	file, err := os.Open(conf)
	if err != nil {
		log.Fatal("open configuration file error:", err)
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("read configuration file error:", err)
	}
	err = json.Unmarshal(b, &Config)
	if err != nil {
		log.Fatal("parse configuration file error:", err)
	}
}
