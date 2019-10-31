package push

import (
	"audience/conf"
	"fmt"
	"net/http"
)

func ServerChan(server string) {
	for k := range conf.Config.ServerChanSCK {
		url := `https://sc.ftqq.com/` + conf.Config.ServerChanSCK[k] + `.send?desp=服务器:` + server + ` 发生故障&text=服务器故障`
		_, err := http.Get(url)
		if err != nil {
			fmt.Println("push error:", err)
		}
	}
}
