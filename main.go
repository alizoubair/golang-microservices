package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/alizoubair/golang-microservices/client"
)

func main() {
	client := client.New("http://localhost:3000")

	price, err := client.FetchPrice(context.Background(), "ETH")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", price)
return
	listenAddr := flag.String("listenaddr", ":3000", "listen address the service is  running")
	flag.Parse()

	svc := NewLoggingService(&priceFetcher{})

	server := NewJSONAPIServer(*listenAddr, svc)
	server.Run()
}