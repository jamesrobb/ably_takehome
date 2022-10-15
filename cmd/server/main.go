package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	"github.com/jamesrobb/ably-takehome/protocol/generated/protocol"
)

func main() {
	port := flag.Int("port", 50051, "port to run server on")
	_ = port
	flag.Parse()

	go startNumberServer(*port)
	fmt.Println("listening...")

	waitForTerminationSignal()
}

func startNumberServer(port int) {
	var opts []grpc.ServerOption

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		fmt.Printf("unable to create TCP listener: %s\n", err)
		os.Exit(1)
	}

	grpcServer := grpc.NewServer(opts...)
	inMemoryStorage := NewInMemoryStorage()
	ns := newNumberServer(inMemoryStorage)
	protocol.RegisterNumbersServer(grpcServer, ns)
	err = grpcServer.Serve(lis)
	if err != nil {
		fmt.Printf("unable to start gRPC server: %s\n", err)
		os.Exit(1)
	}
}

func waitForTerminationSignal() {
	osSignal := make(chan os.Signal, 1)

	signal.Notify(osSignal, syscall.SIGINT)
	signal.Notify(osSignal, syscall.SIGTERM)

	<-osSignal
}
