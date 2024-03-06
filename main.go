package main

import (
	loglib "bitbucket.bri.co.id/nds/nds-go-module-logger/lib"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"quiz-nds-be/client"
	"quiz-nds-be/controller"
	"quiz-nds-be/generated/proto"
	"quiz-nds-be/service"
	"syscall"
)

func main() {
	logger := loglib.NewLib()
	logger.Init("NDS", "V1.0.0")
	cl := client.NewClientStruct(logger)
	ctr := controller.NewController(logger, cl)
	svc := service.NewService(logger, ctr)

	// age
	clAge := client.NewClientAgeStruct(logger)
	ctrAge := controller.NewControllerAge(logger, clAge)
	svcAge := service.NewServiceAge(logger, ctrAge)

	grpcServer := grpc.NewServer()

	proto.RegisterBangunDatarServer(grpcServer, svc)
	proto.RegisterAgeServer(grpcServer, svcAge)
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", "", 8081))
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to start Ln Service gRPC server: %v", err)
		}
	}()

	fmt.Printf("gRPC server is running at %s:%d\n", "", 8081)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	signal := <-c
	log.Fatalf("process killed with signal: %v\n", signal.String())
}
