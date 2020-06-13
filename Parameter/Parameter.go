package Parameter

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func Get_Parameters(ips []string) map[string][]float64 {

	Table := make(map[string][]float64)


	for i :=0 ;i<len(ips); i++ {

		resp, err := http.Get("http://"+ips[i]+":9999/")
		if err!= nil {
			log.Fatalln(err)
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
				Table[k] = append(Table[k], vaa)
			default :
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
