package check

import (
    "fmt"
    "net"
)

func TcpCheck(url string)  {
    _, err := net.Dial("tcp", url)
    if err != nil {
        fmt.Println("tcp check error:",err)
    }
}
