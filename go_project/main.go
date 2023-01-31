package main

import (
	"fmt"
	"net"
	"os"
	projectping "project/project_ping"
	"time"
)

func main() {

	// if len(os.Args) < 2 {
	// 	usage()
	// }

	host := os.Args[1]
	raddr, err := net.ResolveIPAddr("ip", host)
	// fmt.Println("ipppp",raddr)
	if err != nil {
		fmt.Printf("Fail to resolve %s, %s\n", host, err)
		return
	}

	fmt.Printf("Ping %s (%s):\n\n", raddr.String(), host)

	for i := 1; i < 6; i++ {
		if err = projectping.SendICMPRequest(projectping.GetICMP(uint16(i)), raddr); err != nil {
			fmt.Printf("Error: %s\n", err)
		}
		time.Sleep(2 * time.Second)
	}
}
