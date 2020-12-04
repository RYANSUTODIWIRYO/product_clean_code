package api

import (
	//"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	productService "github.com/product/pkg/adapter/api/grpc"
	productSvc "github.com/product/pkg/infrastructure/grpc/proto/product"
	ucProduct "github.com/product/pkg/usecase/product"
	ucCreateProduct "github.com/product/pkg/usecase/create_product"

	container "github.com/product/pkg/shared/di"

	"google.golang.org/grpc"
)

type validator interface {
	Validate() error
}

// RunServer :
func RunServer() {
	log.Println("Starting gRPC Server...")
	//config := cfg.GetConfig()

	//tls := config.Server.Grpc.TLS
	opts := []grpc.ServerOption{}
	//if tls {
	//	serverCert := "server.crt"
	//	serverKey := "server.key"
	//	creds, err := credentials.NewServerTLSFromFile(serverCert, serverKey)
	//	if err != nil {
	//		log.Fatalln("Failed to loading certificates: ", err)
	//		os.Exit(1)
	//	}
	//
	//	opts = append(opts, grpc.Creds(creds))
	//}
	//opts = append(opts, grpc.UnaryInterceptor(serverInterceptor))

	grpcServer := grpc.NewServer(opts...)
	ctn := container.NewContainer()

	Apply(grpcServer, ctn)

	svcHost := "localhost"
	svcPort := 8080

	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", svcHost, svcPort))
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to start BL gRPC server: %v", err)
		}
	}()

	fmt.Printf("Clean Code gRPC server is running at %s:%d\n", svcHost, svcPort)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	signal := <-c
	log.Fatalf("process killed with signal: %v\n", signal.String())
}

// Apply :
func Apply(server *grpc.Server, ctn *container.Container) {
	productSvc.RegisterProductServiceServer(server, productService.NewProductService(ctn.Resolve("ProductSvc").(*ucProduct.ProductInteractor)))
	productSvc.RegisterCreateProductServiceServer(server, productService.NewCreateProductService(ctn.Resolve("CreateProductSvc").(*ucCreateProduct.CreateProductInteractor)))
}
