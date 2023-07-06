package main

import (
	"github.com/akonovalovdev/servers/grpc/grpcAdder/pkg/adder"
	api "github.com/akonovalovdev/servers/grpc/grpcAdder/pkg/api/genproto/adder"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// 1 шаг, создаём сам непосредственно сервер(узел)
	s := grpc.NewServer()
	// структура, которая реализует интерфейс нашего jrpc сервера
	srv := &adder.GRPCServer{}
	// регистрируем созданный jrpc сервер srv в качестве сервера s, с помощью сгенерированного метода в протобаффере
	api.RegisterAdderServer(s, srv)
	// после запуска сервера выводим в консоль инфу что сервер зарущен
	log.Print("Сервер запущен")
	// создаём слушателя (лиснера) который будет принимать запросы
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal()
	}
	// вызваем у нашего сервера метод serve, который активирует слушателя
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}

// для подключения в отдельном терминале запускаем evans api/proto/adder.proto -p 8080
// для запуска call add
