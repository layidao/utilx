package utilx

import (
	"net"
	"os"
	"strings"
)

func HostName() string {
	host, _ := os.Hostname()
	return host
}

func ServerIPString() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return ""
	}

	ips := make([]string, 0)
	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP.String())
			}

		}
	}

	return strings.Join(ips, ",")
}

func ServerIP() (string, string) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return "", ""
	}
	var internalIP, externalIP string
	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipaddr := ipnet.IP.String()
				if strings.HasPrefix(ipaddr, "10.") || strings.HasPrefix(ipaddr, "192.") || strings.HasPrefix(ipaddr, "127.") {
					internalIP = ipaddr
				} else {
					externalIP = ipaddr
				}
			}

		}
	}

	return externalIP, internalIP
}

func SerIP() string {
	eIP, iIP := ServerIP()
	if eIP != "" {
		return eIP
	} else {
		return iIP
	}
}
