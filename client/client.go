package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	gRPC "github.com/ChrBank/DISYSExercise5/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var clientsName = flag.String("name", "default", "Senders name")
var serverPort = flag.String("server", "5400", "Tcp server")

var server gRPC.TimeSyncClient  //the server
var ServerConn *grpc.ClientConn //the server connection

func main() {
	flag.Parse()

	fmt.Println("--- CLIENT APP ---")

	fmt.Println("--- join Server ---")

	ConnectToServer()
	defer ServerConn.Close()

	parseInput()

}

func parseInput() {
	panic("unimplemented")
}

func ConnectToServer() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))

	//use context for timeout on the connection
	timeContext, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel() //cancel the connection when we are done

	log.Printf("client %s: Attempts to dial on port %s\n", *clientsName, *serverPort)
	conn, err := grpc.DialContext(timeContext, fmt.Sprintf(":%s", *serverPort), opts...)
	if err != nil {
		log.Printf("Fail to Dial : %v", err)
		return
	}

	server = gRPC.NewTimeSyncClient(conn)
	ServerConn = conn
	log.Println("the connection is: ", conn.GetState().String())
}
