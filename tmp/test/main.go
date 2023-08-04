package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

var (
	NetworkInterfaces map[string]net.Interface
)

func availableInterfaces() {

	interfaces, err := net.Interfaces()

	log.Printf("interfaces:%v", interfaces)

	if err != nil {
		fmt.Print(err)
		os.Exit(0)
	}

	fmt.Println("Available network interfaces on this machine : ")
	for _, i := range interfaces {
		// fmt.Printf("Name : %v, interface:%v\n", i.Name, i)
		byNameInterface, err := net.InterfaceByName(i.Name)
		if err != nil {
			fmt.Printf("err:%s", err.Error())
			os.Exit(0)
		}
		// fmt.Printf("interface detail:%v\n", byNameInterface)
		fmt.Printf("%s, %v\n", i.Name, byNameInterface.Flags.String())
	}
}

// GetNetworkInterfaces get all net interfaces
func GetNetworkInterfaces() (map[string]net.Interface, error) {
	NetworkInterfaces = make(map[string]net.Interface)
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	// fmt.Println("interfaces:", interfaces)
	for _, i := range interfaces {
		byNameInterface, err := net.InterfaceByName(i.Name)
		if err != nil {
			return nil, err
		}
		fmt.Println("byNameInterface", byNameInterface)
		NetworkInterfaces[i.Name] = *byNameInterface
		// fmt.Printf("%s, %v\n", i.Name, byNameInterface.Flags.String())
	}
	return NetworkInterfaces, nil
}

// GetNetworkInterfaceFlagUp get net interface up/down status
func GetNetworkInterfaceFlagUp(iface string) string {
	byNameInterface, err := net.InterfaceByName(iface)
	if err != nil {
		return "down"
	}
	if strings.Contains(byNameInterface.Flags.String(), "up") {
		return "up"
	} else {
		return "down"
	}
}

func main() {

	// if len(os.Args) != 2 {
	// 	fmt.Printf("Usage : %s <interface name>\n", os.Args[0])
	// 	os.Exit(0)
	// }

	// ifName := os.Args[1]

	// byNameInterface, err := net.InterfaceByName(ifName)

	// if err != nil {
	// 	fmt.Println(err, "["+ifName+"]")
	// 	fmt.Println("-----------------------------")
	// 	availableInterfaces()
	// 	os.Exit(0)
	// }

	// if strings.Contains(byNameInterface.Flags.String(), "up") {
	// 	fmt.Println("Status : UP")
	// } else {
	// 	fmt.Println("Status : DOWN")
	// }
	// availableInterfaces()

	_, err := GetNetworkInterfaces()
	if err != nil {
		fmt.Println(err.Error())
	}
	// fmt.Println(NetworkInterfaces)
	fmt.Println(GetNetworkInterfaceFlagUp("awdl01"))

}
