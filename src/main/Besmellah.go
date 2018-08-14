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
    
    // for _, file := range fm.Search("/home/mym", "Main.java", true){
    //     fmt.Println(file)
    // }

    for {
		if hwnd := getWindow("GetForegroundWindow") ; hwnd != 0 {
			text := GetWindowText(HWND(hwnd))
			fmt.Println("window :", text, "# hwnd:", hwnd)
		}
	}
}