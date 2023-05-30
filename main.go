package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"monster-db/monster"
	"net"
	"time"

	"github.com/DataDog/datadog-go/v5/statsd"
	//log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
)

type myMonsterServer struct {
}

func (s myMonsterServer) GetMonster(context.Context, *monster.MonsterName) (*monster.Monster, error) {
	return &monster.Monster{
		Name:  &monster.MonsterName{Name: "Orc"},
		Level: 2,
	}, nil
}

func main() {
	statsd, err := statsd.New("127.0.0.1:8125")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("hellow world")
	log.Print("Server started")

	var i float64
	for {
		i += 1
		statsd.Set("example_metric.set", fmt.Sprintf("%f", i), []string{"environment:dev"}, 1)
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	}

	lis, err := net.Listen("tcp", ":8081")

	if err != nil {
		log.Fatalf("cannot create listner: %s", err)
	}

	serverRegistrar := grpc.NewServer()
	service := &myMonsterServer{}

	monster.RegisterMonsterServiceServer(serverRegistrar, service)

	fmt.Println("Starting server")
	err = serverRegistrar.Serve(lis)

	if err != nil {
		log.Fatalf("cannot serve: %s", err)
	}

	statsd.Incr("server_started.increment", []string{"environment:local"}, 1)

}
