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


	//var Table_Parameter map[string][]float64
    var port int64
	ips       := Ip.Get_ips(&port)

	weights   := Weights.Get_Weights()
    Threshold := Weights.Get_Threshold()

    for{
		Table_Parameter := Parameter.Get_Parameters(ips)

		answer := Algo.Algorithm(Table_Parameter,weights,Threshold,ips)


		fmt.Println(ips[answer])
		time.Sleep( 1 * time.Second )
	}




}