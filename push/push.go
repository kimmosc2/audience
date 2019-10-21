package push

import (
    "audience/conf"
    "fmt"
    "net/http"
)

func ServerChan(server string) {
    url :=`https://sc.ftqq.com/`+ conf.Config.ServerChanSCK +`.send?desp=服务器:`+ server +` 发生故障&text=服务器故障`
    _, err := http.Get(url)
    if err != nil {
        fmt.Println("error:",err)
    }
}
