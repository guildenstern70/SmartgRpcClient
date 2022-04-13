/*
 * Project gRpc Client
 * Copyright (c) Alessio Saltarin 2022.
 * Licensed under MIT Licence - See LICENSE
 */

package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/guildenstern70/grpcclient/grpcclient"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

// VERSION OF gRpc Client
const VERSION = "0.1"

const (
	defaultName = "grpc-client"
)

var (
	addr = flag.String("addr", "localhost:9000", "the address to connect to")
)

func main() {
	flag.Parse()
	fmt.Printf("gRpc Client %s", VERSION)

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("cannot close grpc connection: %s", err)
		}
	}(conn)

	client := pb.NewMd5GrpcClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	request := pb.Md5Request{
		StringToHash: "AlessioSaltarin",
	}
	response, err := client.Md5Service(ctx, &request)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Hashed Message: %s", response.HashCode)
}
