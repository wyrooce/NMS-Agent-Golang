package main

import "fmt"
//import "runtime"
import "os"
import "net"

func main(){
	
	host, _ := os.Hostname()
    addrs, _ := net.LookupIP(host)
    for _, addr := range addrs {
        if ipv4 := addr.To4(); ipv4 != nil {
        fmt.Println("IPv4: ", ipv4, host)
    }   
}

}