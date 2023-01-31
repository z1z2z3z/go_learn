/*
 * @Description: ping的实现
 */
package projectping

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

type ICMP struct {
	Type        uint8
	Code        uint8
	CheckSum    uint16
	Identifier  uint16
	SequenceNum uint16
}

func GetICMP(seq uint16) ICMP {
	icmp := ICMP{
		Type:        8,
		Code:        0,
		CheckSum:    0,
		Identifier:  0,
		SequenceNum: seq,
	}

	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, icmp)
	icmp.CheckSum = CheckSum(buffer.Bytes())
	buffer.Reset()

	return icmp
}

func CheckSum(data []byte) uint16 {
	var (
		sum    uint32
		length int = len(data)
		index  int
	)
	for length > 0 {
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		index += 2
		length -= 2
	}
	if length > 0 {
		sum += uint32(data[index])
	}
	sum += (sum >> 16)
	return uint16(^sum)
}

func SendICMPRequest(icmp ICMP, destAddr *net.IPAddr) error {
	// 创建ICMP报文
	conn, err := net.DialIP("ip4:icmp", nil, destAddr)
	if err != nil {
		fmt.Printf("Fail to connect to remove host: %s\n", err)
		return err
	}
	defer conn.Close()

	// 填充报文并发送
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, icmp)

	if _, err := conn.Write(buffer.Bytes()); err != nil {
		return err
	}

	tStart := time.Now()

	conn.SetReadDeadline((time.Now().Add(time.Second * 2)))

	// 接收请求
	recv := make([]byte, 1024)
	receiveCnt, err := conn.Read(recv)

	if err != nil {
		return err
	}

	tEnd := time.Now()
	duration := tEnd.Sub(tStart).Nanoseconds() / 1e6

	fmt.Printf("%d bytes from %s: seq=%d time=%dms\n", receiveCnt, destAddr.String(), icmp.SequenceNum, duration)

	return err
}

////  package main  才可以go build   ||    go run

// func main() {
// 	// if len(os.Args) < 2 {
// 	// 	usage()
// 	// }

// 	host := os.Args[1]
// 	raddr, err := net.ResolveIPAddr("ip", host)
// 	if err != nil {
// 		fmt.Printf("Fail to resolve %s, %s\n", host, err)
// 		return
// 	}

// 	fmt.Printf("Ping %s (%s):\n\n", raddr.String(), host)

// 	for i := 1; i < 6; i++ {
// 		if err = SendICMPRequest(GetICMP(uint16(i)), raddr); err != nil {
// 			fmt.Printf("Error: %s\n", err)
// 		}
// 		time.Sleep(2 * time.Second)
// 	}
// }
