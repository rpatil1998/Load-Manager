package Ip

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"client/Notification"
)

func Get_ips(port *string) []string{

	jsonFile, _ := os.Open("serverDetail.json")
	bytes, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()

	var f interface{}

	json.Unmarshal(bytes,&f)

	m := f.(map[string]interface{})
	ip := make([]string, 0)
	for _ , v := range m {
		switch vv := v.(type) {
		case string :                  // String Ip
			ip = append(ip, vv)
        default :                      //  Raise a Notification that
			i:=fmt.Sprintf("%v", v)
            Notification.Notify_NotStringIp(i)
        	continue
		}
	}
	*port = ip[len(ip)-1]       // last element in config in port number

	ip = ip[:len(ip)]           // remove last element (port)
	return ip
}

