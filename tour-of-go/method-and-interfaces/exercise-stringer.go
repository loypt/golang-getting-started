package main

import (
    "fmt"
    "strings"
    "strconv"
)

type IPAddr [4]byte

func (ipadr IPAddr) String() string{
    s := make([]string, len(ipadr))
    for i, val := range ipadr{
        s[i] = strconv.Itoa(int(val))
    }
    return strings.Join(s, ".")
}

func main(){
    hosts := map[string]IPAddr{
        "loopback": {127, 0, 0, 1},
        "googleDNS": {8, 8, 8, 8},
    }
    for name, ip := range hosts{
        fmt.Printf("%v: %v\n", name, ip)
    }
}
