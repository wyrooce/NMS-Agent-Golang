package main

import "fmt"
//import "runtime"
import "../plugins"
import "log"

func main(){
	
    si := plugins.SystemInfo{}
    ip := si.GetIPAddress()
    fmt.Println(ip)//chand ip nadarim?
    fmt.Println(si.GetHostname())

    as, err := si.GetMACAddress()
    if err != nil {
        log.Fatal(err)
    }
    for _, a := range as {
        fmt.Println(a)
    }

    

    

}