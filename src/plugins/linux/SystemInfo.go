package plugins

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"strconv"

	"golang.org/x/sys/windows/registry"
    "log"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	nett "github.com/shirou/gopsutil/net"
	"github.com/matishsiao/goInfo"
)

//Interface type use for some funcition
type Interface struct{
	name string
	MAC string
	IPs []string
}

type SystemInfo struct {
	systemName    string
	systemCode    string
	ip            string
	//mac           string
	networkDomain string
	user          string
	state         string //on or off
	osVersion     string
	osArch        string
	Interfaces 	  []Interface
}

func (si *SystemInfo) GetSystemCode() string {
	return "systemCode"
}

func (si *SystemInfo) GetIPAddress() string {
	host, _ := os.Hostname()
	// fmt.Println("IPs------------------------")
	// x, _ := net.InterfaceAddrs()
	// fmt.Println(x)
	// fmt.Println("---------------------------")
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			return ipv4.String()
		}
	}
	return "IP Not Found"
}


//GetHostname => hostname
func (si *SystemInfo) GetHostname() string {
	host, _ := os.Hostname()
	si.systemName = host
	return si.systemName
}

// OSName return OS name: windows or linux
func (si *SystemInfo) OSName() string {
	return runtime.GOOS
}

// TotalMemory calc memory capaciy
func (si *SystemInfo) TotalMemory() string {
	vmStat, err := mem.VirtualMemory()
	dealwithErr(err)
	return strconv.FormatUint(vmStat.Total, 10) // in byte
}

// TotalDisk returns 3 value in string byte
func (si *SystemInfo) TotalDisk() (string, string, string) {
	diskStat, err := disk.Usage("/")
	dealwithErr(err)

	total := strconv.FormatUint(diskStat.Total, 10)
	used := strconv.FormatUint(diskStat.Used, 10)
	free := strconv.FormatUint(diskStat.Free, 10)

	return total, used, free
}

//CPUCoreNo => number of cpu core
func (si *SystemInfo) CPUCoreNo() string {
	cpuStat, err := cpu.Info()
	dealwithErr(err)

	fmt.Println(cpuStat[0].Family)
	return strconv.FormatInt(int64(cpuStat[0].Cores), 10)
}

// CPUModelName => Intel(R) Core(TM) i5-2450M CPU @ 2.50GHz
func (si *SystemInfo) CPUModelName() string {
	cpuStat, err := cpu.Info()
	dealwithErr(err)

	return cpuStat[0].ModelName
}

//OSArch => for example: x86_64
func (si *SystemInfo) OSArch() string {
	gi := goInfo.GetInfo()
	return gi.Platform
}

//OSVersion => window 10
func (si *SystemInfo) OSVersion() string {
k, err := registry.OpenKey(registry.LOCAL_MACHINE, "SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion", registry.QUERY_VALUE)
    if err != nil {
        log.Fatal(err)
    }
    defer k.Close()

    pn , _, err := k.GetStringValue("ProductName")
    if err != nil {
        log.Fatal(err)
	}
	return pn
}


//MACAddress => return all interface macAddress with its name
func (si *SystemInfo) MACAddress() ([]string, []string, error) {
	ifa, err := nett.Interfaces()

	if err != nil{
		return nil, nil, err
	}

	var names []string
	var addr []string
	
		for _, v := range ifa {
			names = append(names, v.Name)
			addr = append(addr, v.HardwareAddr)
		}
		return names, addr, nil
}
	
	

func dealwithErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
