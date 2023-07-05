package main

import (
	"context"
	"flag"
	api "github.com/akonovalovdev/servers/grpc/grpcAdder/pkg/api/genproto/adder"
	"google.golang.org/grpc"
	"log"
	"strconv"
)

func main() {
	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatal("not enough arguments")
	}

	x, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := api.NewAdderClient(conn)
	res, err := c.Add(context.Background(), &api.AddRequest{X: int32(x), Y: int32(y)})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.GetResult())

}
