package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"

	"maizuo.com/back-end/iris-demo/src/server/proto/mallorderproto"
	"maizuo.com/back-end/iris-demo/src/server/proto/mallproductproto"

)

func SetupRPC() {

	order_port := viper.GetString("server.order.port")
	order_host := viper.GetString("server.order.host")
	order_address := order_host + order_port
	fmt.Println(order_address)
	// Set up a connection to the server.

	order_conn, err := grpc.Dial(order_address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	orderClient := mallorderproto.NewOrderServiceClient(order_conn)
	OrderClient = orderClient

	product_port := viper.GetString("server.product.port")
	product_host := viper.GetString("server.product.host")
	product_address := product_host + product_port
	fmt.Println(product_address)
	// Set up a connection to the server.
	product_conn, err := grpc.Dial(product_address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	productClient := mallproductproto.NewProductServiceClient(product_conn)
	ProductClient = productClient
}

var (
	OrderClient   mallorderproto.OrderServiceClient
	ProductClient mallproductproto.ProductServiceClient
)
