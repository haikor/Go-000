package main

import (
	"bufio"
	"log"
	"net"
	"strings"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:5000")
	if checkError(err, "start error") {
		return
	}

	log.Println("listen localhost:5000\npress q to quit")
	for {
		conn, err := listen.Accept()
		if checkError(err, "accept error") {
			return
		}
		bytes := make(chan string)
		go parse(conn, bytes)

		go res(conn, bytes)
	}

}

func res(conn net.Conn, bytes chan string) {

	writer := bufio.NewWriter(conn)
	for {
		str := <-bytes
		if strings.Trim(str, "\r\n") == "q" {
			log.Println("exit ")
			err := writer.Flush()
			if checkError(err, "flush error") {
				return
			}
			err = conn.Close()
			if checkError(err, "close error") {
				return
			}
			return
		}
		_, err := writer.WriteString("response:" + str)
		log.Println("write ", str)
		if checkError(err, "write error") {
			return
		}
		err = writer.Flush()
		if checkError(err, "flush error") {
			return
		}
	}

}

func checkError(err error, msg string) bool {
	if err != nil {
		log.Println(msg, err.Error())
		return true
	}
	return false
}

func parse(conn net.Conn, bytes chan string) {
	reader := bufio.NewReader(conn)
	for {
		readString, err := reader.ReadString('\n')

		if checkError(err, "read error") {
			return
		}
		log.Println("receive ", readString)

		bytes <- readString
		if readString == "q" {
			return
		}
	}

}
