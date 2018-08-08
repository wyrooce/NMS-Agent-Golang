package plugins

import (
	"net"
	"os"
)

type SystemInfo struct{
	systemName string
	systemCode string
	ip string
	mac string
	networkDomain string
	user string
	state string //on or off
	osVersion string
	osArch string
}

func (si *SystemInfo) GetSystemCode() string{
	return "systemCode"
}

func (si *SystemInfo) GetIPAddress() string{
	host, _ := os.Hostname()
    addrs, _ := net.LookupIP(host)
    for _, addr := range addrs {
        if ipv4 := addr.To4(); ipv4 != nil {
			return ipv4.String()
    	}   
	}
	return "IP Not Found"
}

func (si *SystemInfo) GetMACAddress() ([]string, error) {
    iface, err := net.Interfaces()
    if err != nil {
        return nil, err
    }
    var as []string
    for _, ifa := range iface {
        a := ifa.HardwareAddr.String()
        if a != "" {
            as = append(as, a)
        }
    }
    return as, nil
}


func (si *SystemInfo) GetHostname() string{
	host, _ := os.Hostname()
	si.systemName = host
    return si.systemName
}


