package network

import "net"



func Listen(port int, handler func(data [1500]uint8, rlen int)) {
	conn,err := net.ListenUDP("udp", &net.UDPAddr{Port: port})
	if err != nil {
		panic(err)
	}
	var buff [1500]uint8
	for {
		rlen, _, err := conn.ReadFromUDP(buff[:])
		if err != nil {
			panic(err)
		}
		handler(buff, rlen)

	}
}
