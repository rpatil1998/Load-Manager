package Ip

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"client/Notification"
)

func Get_ips(port *int64) []string{

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
		case int64 :                  // Port Number
			*port = vv
        default :                      //  Raise a Notification that
            Notification.Notification_wrongIp(string(v))
        	continue
		}
	}
	return ip
}

