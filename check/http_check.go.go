package check

import (
    "audience/push"
    "fmt"
    "net/http"
    "time"
)

// checkSilent is a simple function for check, checkSilent
// no logs.
func checkSilent(url string, t time.Duration) {
    for {
        if _, err := http.Get(url); err != nil {
            push.ServerChan(url)
        }
        time.Sleep(t)
    }
}

// Default check function, if target server no response,
// check will call ServerChan function to notice human.
// check has log out.
func check(url string, t time.Duration) {
    for {
        if _, err := http.Get(url); err != nil {
            fmt.Println("request error:", err)
            push.ServerChan(url)
        } else {
            fmt.Printf("the %v check pass\n", url)
        }
        time.Sleep(t)
    }
}
