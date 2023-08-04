package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func availableInterfaces() {

	interfaces, err := net.Interfaces()

	if err != nil {
		fmt.Print(err)
		os.Exit(0)
	}

	fmt.Println("Available network interfaces on this machine : ")
	for _, i := range interfaces {
		fmt.Printf("Name : %v \n", i.Name)
	}
}

func main() {

	if len(os.Args) != 2 {
		fmt.Printf("Usage : %s <interface name>\n", os.Args[0])
		os.Exit(0)
	}

	ifName := os.Args[1]

	byNameInterface, err := net.InterfaceByName(ifName)

	if err != nil {
		fmt.Println(err, "["+ifName+"]")
		fmt.Println("-----------------------------")
		availableInterfaces()
		os.Exit(0)
	}

	if strings.Contains(byNameInterface.Flags.String(), "up") {
		fmt.Println("Status : UP")
	} else {
		fmt.Println("Status : DOWN")
	}

}
