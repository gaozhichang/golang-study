/*
@Time : 2021/8/31 下午8:55
@Author : gaozhichang
@File : http.go
@des: GoLand
*/

package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
)

func main() {
	log.SetFlags(log.LstdFlags|log.Lshortfile)
	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Panic(err)
	}

	//c:=make(chan int)
	//go func(c chan  int) {
	//	time.Sleep(10 * time.Second)
	//	c<-1
	//}(c)

	for {
		log.Println("start..",l)
		client, err := l.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handleHttpRequest(client)
	}
}

func handleHttpRequest(client net.Conn) {
	if client == nil {
		return
	}
	defer client.Close()

	var b [1024]byte
	_, err := client.Read(b[:])
	if err != nil {
		log.Println(err)
		return
	}
	var method, host string
	fmt.Sscanf(string(b[:bytes.IndexByte(b[:], '\n')]), "%s%s", &method, &host)
	//hostPortURL, err := url.Parse(host)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//log.Println(host)
	//log.Println(method)
	//log.Println(address)
	//log.Println(hostPortURL)


	fmt.Fprint(client, "HTTP/1.1 200 OK\r\n")
	fmt.Fprint(client, "COOKIE: name=gao")
	fmt.Fprint(client, "\n\n")
	fmt.Fprint(client, fmt.Sprintf("%s %s\n%s",method,host,"hello world!"))
	log.Println("done")

}
