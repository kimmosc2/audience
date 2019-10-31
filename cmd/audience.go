package main

import (
	"audience/check"
	"audience/conf"
)

func main() {
	for k := range conf.Config.Server {
		if conf.Config.LogMode {
			go check.HttpCheck(conf.Config.Server[k])
		} else {
			go check.HttpCheckSilent(conf.Config.Server[k])
		}
	}
	ch := make(chan int, 0)
	ch <- 0
}
