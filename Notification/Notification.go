package Notification

import "fmt"

func Notify_NotStringIp(ip string){
	fmt.Println(ip +"  is not a string")
	fmt.Println("Please give string in server details configuration")
}

func Notify_ServerNotResponding(ip string)  {

	fmt.Println(ip + " is Not responding")
	fmt.Println(ip+" has been discarded")

}

func Notify_ServerGivingUnexpectedOutput(ip string){

	fmt.Println(ip + "  This server is giving unexpected output")

}

func Relative_percentage_config()  {

	fmt.Println("Please check Relative Percentage Config file")

}
