package main

import (
	"context"
	"fmt"
	"log"
	"monster-db/monster"
	"net"

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
	fmt.Println("hellow world")

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

}
