package conf

import (
    "encoding/json"
    "flag"
    "io/ioutil"
    "log"
    "os"
)

var (
    conf string
)

type configuration struct {
    Server []map[string]int64 `json:"httpServer"`
    ServerChanSCK string `json:"serverChanSCK"`
    LogOut bool `json:"logOut"`
}

var Config configuration

func init() {
    flag.StringVar(&conf,"conf","","configuration file")
    flag.Parse()
    if conf == "" {
        log.Fatal("no configuration file")
    }
    file, err := os.Open(conf)
    if err != nil {
        log.Fatal("open configuration file error:",err)
    }
    defer file.Close()
    b, err := ioutil.ReadAll(file)
    if err != nil {
        log.Fatal("read configuration file error:",err)
    }
    err = json.Unmarshal(b, &Config)
    if err != nil {
        log.Fatal("parse configuration file error:",err)
    }
}

