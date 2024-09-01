package main

import (
	"log"
	"net"
	"io"
)

// Change this variable before building
const TARGET = "httpforever.com:80"
// Connection method
const METHOD = "tcp"

// Local IP
const IP = "127.0.0.1"
// Local port
const PORT = "1234"

func handler(src net.Conn){
	log.Printf("success!")
	dest, err := net.Dial(METHOD, TARGET)
	if err != nil {
		log.Println(err)
		log.Fatalln("proxy destination connection error")
	} 
	defer src.Close()
	go func(){
		if _, err := io.Copy(dest, src); err != nil {
			log.Println(err)
			log.Fatalln("unable to copy stream [dest->source]")
		}
	}()
	if _, err := io.Copy(src, dest); err != nil {
		log.Println(err)
		log.Fatalln("unable to copy stream [source->dest]")
	}
}
func listener(){
	log.Printf("listening on port %s", PORT)
	lst, err := net.Listen(METHOD, IP + ":" + PORT)
	if err != nil {
		log.Println(err)
		log.Fatalln("listener bind error; try changing the port?")
	}
	for {
		conn, err := lst.Accept()
		log.Printf("got a connection: %s", conn.RemoteAddr())
		if err != nil {
			log.Println(err)
			log.Fatalln("source connection error")
		}
		go handler(conn)
	}
}
func main(){
	log.Printf("proxying from %s:%s to %s using %s", IP, PORT, TARGET, METHOD)
	listener()
}