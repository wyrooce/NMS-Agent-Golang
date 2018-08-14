package main

import "fmt"
//import "runtime"
import "../plugins"

func main(){
	
    si := plugins.SystemInfo{}
    fm := plugins.FileManager{}
    ip := si.GetIPAddress()
    fmt.Println(ip)//chand ip nadarim?
    fmt.Println(si.GetHostname())
    
    for _, file := range fm.GetFilesInPathRec("/home/mym/IdeaProjects"){
        fmt.Println(file)
    }
}