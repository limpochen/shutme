package cmds

import (
	"bytes"
	"encoding/binary"
	"log"
	"net"
	"os"
	"time"
)

const (
	hostUnknown = 0
	hostOnline  = 1
	hostOffline = 2
)

const ICMP_DATA = 76

var pingCount uint16

// ICMP header
type ICMP struct {
	Type        uint8
	Code        uint8
	Checksum    uint16
	Identifier  uint16
	SequenceNum uint16
}

type PingPacket struct {
	icmp ICMP
	data [ICMP_DATA]byte
}

func init() {
	pingCount = 0
}

func CheckSum(data []byte) (rt uint16) {
	var (
		sum    uint32
		length int = len(data)
		index  int
	)
	for length > 1 {
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		index += 2
		length -= 2
	}
	if length > 0 {
		sum += uint32(data[index]) << 8
	}
	rt = uint16(sum) + uint16(sum>>16)

	return ^rt
}

// Send ICMP packets to the target host and receive the response packets.
// Param:	hostname
// Return:	duration in milliseconds; error message
func Ping(domain string) (float64, error) {
	var (
		pingPacket PingPacket
		laddr      = net.IPAddr{IP: net.ParseIP("0.0.0.0")} // Get localhost IP address structure.
		raddr, _   = net.ResolveIPAddr("ip", domain)        // Get the remote IP address structure by resolving the domain name.
		err        error
		conn       *net.IPConn
	)

	pingCount++

	// Return ip socket
	conn, err = net.DialIP("ip4:icmp", &laddr, raddr)
	if err != nil {
		return -1, err
	}

	defer conn.Close()
	// Initialize ICMP packet
	pingPacket.icmp = ICMP{8, 0, 0, uint16(os.Getpid()), pingCount}

	for i := 0; i < ICMP_DATA; i++ {
		pingPacket.data[i] = byte(i + 1)
	}

	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, pingPacket)
	b := buffer.Bytes()
	binary.BigEndian.PutUint16(b[2:], CheckSum(b))

	recv := make([]byte, 1024)

	//	Send binary packet to the target address.
	if _, err := conn.Write(buffer.Bytes()); err != nil {
		return -1, err
	}

	// Otherwise, record the current time.
	t_start := time.Now()
	conn.SetReadDeadline((time.Now().Add(time.Second * 3)))

	_, err = conn.Read(recv)
	//	Check if the target address fails upon return.
	if err != nil {
		return -1, err
	}
	t_end := time.Now()

	return float64(t_end.Sub(t_start).Nanoseconds()) / 1e6, nil
}

// Continuously probe remote host. Use flags: Flag_t; Flag_n; Flag_i
// Param : none
// Return: none
func ProbeRemote() {
	iOnline := hostUnknown // The remote host status, hostOnline; hostOffline; hostUnknown
	iFailed := 0           // Numbers offline is detected

	//fmt.Println("Probing the host:", Flag_t)
	//MyLog(Info, "Start probing the remote host: "+Flag_t)
	log.Println("Start probing the remote host: " + Flag_t)

	for {
		_, err := Ping(Flag_t)
		if err != nil { // 服务模式：记着处理好错误返回方式
			iFailed++
			if iFailed <= Flag_n {
				if iFailed == 0 { // Detected offline for the first time
					//MyLog(Warning, "Detected remote host offline, possibly due to a power outage.")
					log.Println("Detected remote host offline, possibly due to a power outage.")
				}
				log.Printf("The network is disconnected:%d/%d\n", iFailed, Flag_n)
			}
			iOnline = hostOffline
		} else {
			if iOnline != hostOnline {
				log.Println("The network is connected.")
			}
			if iFailed > 0 { //iOnline == hostOffline
				//MyLog(Info, "The network recovered.")
				log.Println("The network recovered.")
			}
			iFailed = 0
			iOnline = hostOnline
		}

		if iFailed == Flag_n {
			err := ShutMeRun()
			if err != nil {
				//MyLog(Error, fmt.Sprintf("ShutMe failed: %v.", err))
				log.Printf("ShutMe failed: %v.\n", err)
			} else {
				log.Println("ShutMe down.")
				//MyLog(Warning, fmt.Sprintf("ShutMe command \"%s\" has been triggered.", ShutMeCmd))
				log.Printf("ShutMe command \"%s\" has been triggered.\n", ShutMeCmd)
			}
		}
		time.Sleep(time.Second * time.Duration(Flag_i))
	}
}
