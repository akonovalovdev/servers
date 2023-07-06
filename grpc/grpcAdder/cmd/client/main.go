package main

// реализуем клиента grpc сервера
import (
	"context"
	"flag"
	"fmt"
	api "github.com/akonovalovdev/servers/grpc/grpcAdder/pkg/api/genproto/adder"
	"google.golang.org/grpc"
	"log"
	"strconv"
)

func main() {
	// аргументы Х и У получаем при запуске клиента в командной строке передаём их
	flag.Parse()
	// проверяем что передано минимум 2 аргумента при запуске программы
	if flag.NArg() < 2 {
		log.Fatal("not enough arguments")
	}
	// теперь нам нужно получить аргументы из командной строки
	// и перевести значения из строк в инты, так как по умолчанию приходят строки
	x, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Fatal(err)
	}
	// проверяем что аргументы записались
	fmt.Println("x =", x, "y =", y)
	// подключаемся к нашему grpc серверу. Создаём СОЕДИНЕНИЕ к серверу
	conn, err := grpc.Dial(":8080", grpc.WithInsecure()) // передаём наш хост и безопасное соединение
	if err != nil {
		log.Fatal(err)
	}
	// у клиента реализованы точно такие же методы, как у интерфейса сервера, поэтому пользуемся
	// тем же сгенерированным протофайлом, что и у сервера
	// Создаём клиент
	c := api.NewAdderClient(conn)
	// У клиента вызываем метод Add (первый аргумент пустышка-контекст.бэкграунд
	// второй аргумент, это сообщение, которое мы описывали в протофайле
	res, err := c.Add(context.Background(), &api.AddRequest{X: int32(x), Y: int32(y)})
	if err != nil {
		log.Fatal(err)
	}
	// Если всё прошло хорошо, то у ответа есть метод достать результат, его напечатаем в командную строку
	log.Println("Result:", res.GetResult())

}
