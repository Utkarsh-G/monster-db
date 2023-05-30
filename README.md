# monster-db

To start the grpc server:
go run main.go

Then you can hit
grpc://localhost:8081
with a request for MonsterService / GetMonsters with a message like

{
    "name": "veniam amet esse enim occaecat"
}
