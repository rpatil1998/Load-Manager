package Parameter

import (
	"client/Notification"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func Get_Parameters(ips *[]string, port string) map[string][]float64 {

	Table := make(map[string][]float64)

	var temp []string

	for i:=0;i<len(*ips);i++{
		temp = append(temp,(*ips)[i])
	}


	for i :=0 ;i<len(temp); i++ {

		resp, err := http.Get("http://"+temp[i]+port)
		if err!= nil {

			// If there is an error
			// Retry after 100 milisecond
			time.Sleep( 100 * time.Microsecond)
			resp, err = http.Get("http://"+temp[i]+port)

			if err!= nil {

				// Server still not responding  Notify that this particular ip is not responding
				Notification.Notify_ServerNotResponding(temp[i])

				// Removing ip which is not responding
				copy((*ips)[i:],(*ips)[i+1:])
				(*ips)[len(*ips)-1]=""
				*ips = (*ips)[:len(*ips)-1]


				// Go to next ip
				continue

				//log.Fatalln(err)
			}
		}
		defer resp.Body.Close()

		body, readErr := ioutil.ReadAll(resp.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}

		var P interface{}


		json.Unmarshal([]byte(body), &P)

		m := P.(map[string]interface{})

		for k, v := range m {
			switch vaa := v.(type) {
			case string :
				t, _ := strconv.Atoi(vaa)
				tt := float64(t)
				Table[k] = append(Table[k],tt)
			case float64 :
				Table[k] = append(Table[k],vaa)
			default :

				Notification.Notify_ServerGivingUnexpectedOutput(temp[i])
				continue
			}
		}
	}

	Max_connections := float64(0)
	Max_Network :=float64(0)


	for i := 0; i < len(Table["Network_Utilization_out"]); i++ {

		if Max_Network < Table["Network_Utilization_out"][i] {
			Max_Network = Table["Network_Utilization_out"][i]
		}
	}
	for i := 0; i < len(Table["Network_Utilization_out"]); i++ {

		Table["Network_Utilization_out"][i] = Table["Network_Utilization_out"][i]/Max_Network*100
	}




	for i := 0; i < len(Table["Connections"]); i++ {

		if Max_connections < Table["Connections"][i] {
			Max_connections = Table["Connections"][i]
		}
	}
	for i := 0; i < len(Table["Connections"]); i++ {
		Table["Connections"][i] = Table["Connections"][i]/Max_connections *100

	}


	return Table

}
