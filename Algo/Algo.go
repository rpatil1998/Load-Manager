package Algo

import "math"

func Discard_ip(Table map[string][]float64, Threshold map[string]float64, ips []string) []bool{

	var ip_status []bool      // ip status --- is this ip available now after discarding

	// initialise all ip are available first

	// ip[0] available
	//ip[1] available
	//ans so on

	for i := 0; i < len(ips); i++ {
		ip_status = append(ip_status, true)
	}

	// make the status of ip to not available (ip_status = false) if it exceed CPU or Memory thredhold
	for p,_ := range Table{

		if (p=="Cpu_Utilization") || (p=="Memory_Utilization"){

			for i:=int64(0);i<int64(len(ips));i++{

				if Table[p][i] < Threshold[p]{
					ip_status[i]=false
				}

			}
		}

	}


	return  ip_status
}


func Give_best_Available(Table map[string][]float64, weights map[string]float64,ip_status []bool)  int {

	Load :=math.MaxFloat64
	var load_ip float64
	answer := 0

	for i:=0;i<len(ip_status);i++{
		if ip_status[i]==true{
			load_ip = 0
			for p,_ :=range Table{
				load_ip = load_ip + Table[p][i] * weights[p]
			}
		    if load_ip<Load{
				Load = load_ip
				answer=i
			}
		}

	}
	return answer

}

// Use this Algo when no server available after discarding on the basis of CPU & Memory threshold
func Give_best_NotAvailable(Table map[string][]float64, weights map[string]float64, ips []string)  int {

	Load :=math.MaxFloat64
	var load_ip float64
	answer := 0

	//minimum load score ip
	for i:=0;i<len(ips);i++ {
		load_ip = 0
		for p, _ := range Table {
			load_ip = load_ip + Table[p][i]*weights[p]
		}
		if load_ip < Load {
			Load = load_ip
			answer = i
		}

	}


	return answer

}

func Algorithm(Table map[string][]float64, weights map[string]float64,Threshold map[string]float64, ips []string) int {


	//Discarding ips which has Cpu_utilization or memory_Utilization more than threshold
	ip_status := Discard_ip(Table, Threshold, ips)


	// After discardin ips on the basis of threshold (CPU & Memory) is there any ip remains available
	is_any_available := false
	for i := 0; i < len(ips); i++ {
		if ip_status[i]{
			is_any_available=true
		}
	}

	// if any server(ip) avaailable after discarding above thredhold ones than use This algo
	if is_any_available{
		return Give_best_Available(Table,weights,ip_status)
	}


    //  If all server(ip) has been discard just consider all ips dont discard anyone
	return Give_best_NotAvailable(Table,weights,ips)

}

