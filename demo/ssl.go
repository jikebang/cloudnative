package main

import (
	"log"
	"net"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hi", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hi"))
	})
	l,err := net.Listen("tcp",":443")
	if err != nil {
		log.Fatal(err)
	}
	err = http.ServeTLS(l,mux,"./wang.dd+4.pem","./wang.dd+4-key.pem")
	log.Fatal(err)
}
