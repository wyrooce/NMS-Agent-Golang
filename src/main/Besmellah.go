package main

import (
	"fmt"
	"../plugins/windows"
)

func main() {
	pr := fmt.Println
	pr(windows.Property())

}
