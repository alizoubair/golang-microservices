package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/alizoubair/price-fetcher/client"
	"github.com/alizoubair/price-fetcher/proto"
)

func main() {
	var (
		jsonAddr = flag.String("json", ":3000", "listen address of the json transport");
		grpcAddr = flag.String("grpc", ":4000", "listen address of the grpc transport")
		svc      = NewLoggingService(&priceFetcher{})
		ctx      = context.Background()
	)
	flag.Parse()

	grpcClient, err := client.NewGRPCClient(":4000")
	if err != nil {
		log.Fatal(err)
	}

	go func () {
		time.Sleep(3 * time.Second)
		for {
			resp, err := grpcClient.FetchPrice(ctx, &proto.PriceRequest{Ticker: "ETH"})
			if err != nil {
				log.Fatal(err)
			}
			
			fmt.Printf("%+v\n", resp)
		}
	}()

	go makeGRPCServerAndRun(*grpcAddr, svc)

	server := NewJSONAPIServer(*jsonAddr, svc)
	server.Run()
}