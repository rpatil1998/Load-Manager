package main

import (
	"client/Parameter"
    "client/Ip"
	"client/Weights"
	"client/Algo"
	"fmt"
	"time"

)

func main(){

    var port string
	ips       := Ip.Get_ips(&port)

	weights   := Weights.Get_Weights()
    Threshold := Weights.Get_Threshold()

    for{
		//var Table_Parameter map[string][]float64  contains all parameter fo all ips
    	//
    	// parameter : "Cpu_Utilisation", "Memory_Utilization" etc.

		Table_Parameter := Parameter.Get_Parameters(&ips,port)

		// Here answer represent index of best ip

		answer := Algo.Algorithm(Table_Parameter,weights,Threshold,ips)


		fmt.Println(ips[answer])
		time.Sleep( 1 * time.Second )
	}




}