package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net"
	"sort"
	"strings"
)

func GetDeviceID() (devid string, err error) {
	var macString string
	var macs []string
	// 获取本机的MAC地址
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error : " + err.Error())
		return
	}
	for _, inter := range interfaces {
		mac := inter.HardwareAddr //获取本机MAC地址
		//fmt.Println("MAC = ", mac)
		if mac.String() != "" {
			macs = append(macs, mac.String())
		}
	}
	sort.Strings(macs)
	for _, v := range macs {
		macString += v
		macString += ","
	}
	//fmt.Println("macString=", macString)
	h := md5.New()
	h.Write([]byte(macString))
	devid = hex.EncodeToString(h.Sum(nil))
	devid = strings.ToUpper(devid)
	// FB5290764EEE488301ABFB10B0A28FD8
	if len(devid) == 32 {
		devid = devid[0:16]
	}
	return
}
