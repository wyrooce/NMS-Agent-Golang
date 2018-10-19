package windows


import (
	"os/exec"
	"bytes"
	"strings"
	"log"
	"os"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/disk"
	"strconv"
)

//Interface shows a NIC 
type Interface struct{
	name string
	MAC string
	IPs []string
}

//SystemInfo overal information about the host
type SystemInfo struct {
	Hostname    string `json:"hostname,omitempty"`
	OSName string `json:"osname,omitempty"`
	OSVersion string `json:"osversion,omitempty"`
	OSArch string `json:"osarch,omitempty"`
	ProductID string `json:"productid,omitempty"`
	SystemManufacturer string `json:"manufacturer,omitempty"`
	CPU string `json:"cpu,omitempty"`
	CoreNo string `json:"coreno,omitempty"`
	MemorySize string `json:"memorysize,omitempty"`
	DiskSize string `json:"totaldisk,omitempty"`
	DiskUsed string `json:"useddisk,omitempty"`
	Domain string `json:"domain,omitempty"`
	User          string `json:"user,omitempty"`
	Interfaces 	  []Interface `json:"interface,omitempty"`
}

func systeminformation() string{
	cmd := exec.Command("systeminfo")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
	return out.String()
}

//Property parse systeminfo output
func Property() SystemInfo {
	output := systeminformation()
	info := SystemInfo{}
	var value string

	arr := strings.Split(output, "\n")
	for _, v := range arr{
		value = strings.TrimSpace(v[strings.Index(v, ":")+1: len(v)])
		if strings.HasPrefix(v, "Domain:") {
			info.Domain = value
		}else if strings.HasPrefix(v, "System Type:"){
			info.OSArch = value
		}else if strings.HasPrefix(v, "OS Version:"){
			info.OSVersion = value
		}else if strings.HasPrefix(v, "Product ID:"){
			info.ProductID = value
		}else if strings.HasPrefix(v, "OS Name:"){
			info.OSName = value
		}else if strings.HasPrefix(v, "System Manufacturer:"){
			info.SystemManufacturer = value
		}
	}

	info.Hostname, _ = os.Hostname()
	//------------------------------
	cpuStat, err := cpu.Info()
	dealwithErr(err)
	info.CPU = cpuStat[0].ModelName
	info.CoreNo = strconv.FormatInt(int64(cpuStat[0].Cores), 10)
	//------------------------------
	vmStat, err := mem.VirtualMemory()
	dealwithErr(err)
	info.MemorySize = strconv.FormatUint(vmStat.Total/1024/1024, 10) // in byte
	//------------------------------
	diskStat, err := disk.Usage("/")
	dealwithErr(err)
	info.DiskSize = strconv.FormatUint(diskStat.Total/1024/1024, 10)
	info.DiskUsed = strconv.FormatUint(diskStat.Used/1024/1024, 10)

	//fmt.Printf("%+#v\n-----------------\n", info)
	// fmt.Println(systeminfo)
	fmt.Println("Property() result returned.")
	return info
}
 
func dealwithErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}