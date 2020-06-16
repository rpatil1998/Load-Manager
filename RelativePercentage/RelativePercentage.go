package RelativePercentage

import (
	"client/Notification"
	"encoding/json"
	"io/ioutil"
	"os"
)

func Relative_Percent_para() []string{

	jsonFile, _ := os.Open("RelativePercentage.json")
	bytes, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()

	var f interface{}

	json.Unmarshal(bytes,&f)

	m := f.(map[string]interface{})
	Relative_percent_parameter := make([]string, 0)
	for _ , v := range m {
		switch vv := v.(type) {
		case string :                  // String Ip
			Relative_percent_parameter = append(Relative_percent_parameter, vv)
		default :                      //  Raise a Notification that
			Notification.Relative_percentage_config()
			continue
		}
	}
	         // remove last element (port)
	return Relative_percent_parameter
}


