package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net"
	"time"
)

const (
	defaultSourceIP   = "116.202.110.84"
	defaultSourcePort = 21820
	defaultDestIP     = "0.0.0.0"
	defaultDestPort   = 56789

	defaultLength        = 16
	defaultMinNumPackets = 1
	defaultMaxNumPackets = 4
	defaultProtocol      = "udp"
)

func main() {
	destAddrPort := flag.String("d", "", "set the destination IP address and port in the format IP:Port")
	srcAddrPort := flag.String("s", "", "set the source IP address and port in the format IP:Port")
	lengthOnePacket := flag.Int("l", 16, "set the length of the packet")
	flag.Parse()
	flag.Uint("n", 4, "set the number of packets")

	if *destAddrPort == "" || *srcAddrPort == "" {
		fmt.Println("required -dest and -src flags")
		return
	}

	destIP, destPort, dstErr := net.SplitHostPort(*destAddrPort)
	if dstErr != nil {
		fmt.Println("invalid -dest flag format")
		return
	}

	srcIP, srcPort, srcErr := net.SplitHostPort(*srcAddrPort)
	if srcErr != nil {
		fmt.Println("invalid -src flag format")
		return
	}

	numPackets := rand.Intn(defaultMaxNumPackets-1) + defaultMinNumPackets

	// Создание UDP адреса и подключение к серверу
	remoteAddr, remoteErr := net.ResolveUDPAddr(defaultProtocol, fmt.Sprintf("%s:%s", destIP, destPort))
	if remoteErr != nil {
		fmt.Println("failed to create remote address:", remoteErr)
		return
	}

	// Создание локального адреса
	localAddr, localErr := net.ResolveUDPAddr(defaultProtocol, fmt.Sprintf("%s:%s", srcIP, srcPort))
	if localErr != nil {
		fmt.Println("failed to create local address:", localErr)
		return
	}
	conn, err := net.DialUDP("udp", localAddr, remoteAddr)
	if err != nil {
		fmt.Println("Ошибка при установлении соединения:", err)
		return
	}
	defer conn.Close()
	//totalBytes := 0
	for i := 0; i < numPackets; i++ {
		// Генерация случайных данных заданной длины
		rand.Seed(time.Now().UnixNano())
		data := make([]byte, *lengthOnePacket)
		rand.Read(data)

		// Отправка данных
		//	totalBytes, err := conn.Write(data)
		if err != nil {
			fmt.Println("Ошибка при отправке данных:", err)
			return
		}

		fmt.Printf("Пакет %d успешно отправлен.\n", i+1)

		// Пауза между отправкой пакетов
		//time.Sleep(1 * time.Second)
	}
}
