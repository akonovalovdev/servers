package main

import (
	"github.com/akonovalovdev/servers/grpc/grpcAdder/pkg/adder"
	api "github.com/akonovalovdev/servers/grpc/grpcAdder/pkg/api/genproto/adder"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// сам непосредственно сервер
	s := grpc.NewServer()
	// структура, которая реализует интерфейс
	srv := &adder.GRPCServer{}
	// регистрируем созданный сервер с помощью сгенерированного метода в протобаффере
	api.RegisterAdderServer(s, srv)
	log.Print("Сервер запущен")
	// создаём слушателя
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal()
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}

// для подключения в отдельном терминале запускаем evans api/proto/adder.proto -p 8080
// для запуска call add
