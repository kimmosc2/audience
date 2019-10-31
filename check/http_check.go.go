package check

import (
	"audience/conf"
	"audience/push"
	"fmt"
	"net/http"
	"time"
)

// HttpCheckSilent is a simple function for check, HttpCheckSilent
// no logs.
func HttpCheckSilent(serverInfo conf.HttpServer) {
	tolerance := 0
	for {
		if _, err := http.Get(serverInfo.Address); err != nil {
			if tolerance >= serverInfo.Tolerance {
				push.ServerChan(serverInfo.Address)
				tolerance = 0
				<-time.NewTimer(serverInfo.PushInterval).C
			} else {
				tolerance++
			}
		}
		time.Sleep(serverInfo.CheckInterval)
	}
}

// Default check function, if target server no response,
// check will call ServerChan function to notice human.
// check has log out.
func HttpCheck(serverInfo conf.HttpServer) {
	tolerance := 0
	for {
		if _, err := http.Get(serverInfo.Address); err != nil {
			fmt.Println("request error:", err)
			if tolerance >= serverInfo.Tolerance {
				push.ServerChan(serverInfo.Address)
				tolerance = 0
				<-time.NewTimer(serverInfo.PushInterval).C
			} else {
				tolerance++
				fmt.Printf("the %v check pass\n", serverInfo.Address)
			}
		}
		time.Sleep(serverInfo.CheckInterval)
	}
}
