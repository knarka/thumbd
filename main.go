package main


import "log"
import "net"

func answer(conn net.Conn) {
	conn.Write([]byte("hello"))
	conn.Close()
}

func main() {
	ln, err := net.Listen("tcp", ":79")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go answer(conn)
	}
}
