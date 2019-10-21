package main

import (
    check "audience/check"
    "audience/conf"
    "time"
)

func main() {
    for _, v := range conf.Config.Server {
        for i, j := range v {
            if conf.Config.LogOut{
                go check.HttpCheck(i, time.Duration(j)*time.Second)
            } else {
                go check.HttpCheckSilent(i, time.Duration(j)*time.Second)
            }
        }
    }
    ch := make(chan int, 0)
    ch <- 0
}

