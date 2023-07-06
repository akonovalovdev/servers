package adder

// файл реализующий в первую очередь интерфейс AdderServer interface
import (
	"context"
	api "github.com/akonovalovdev/servers/grpc/grpcAdder/pkg/api/genproto/adder"
)

// GRPCServer ...
type GRPCServer struct {
	api.UnimplementedAdderServer
}

// Add ...
func (s *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	// мы получаем контекст и request- который описывали в протофайле Х и У
	// то есть протокол баффер нам сгенерировал такие же типы, чтобы всё соответствовало файлам со схемой
	return &api.AddResponse{Result: req.GetX() + req.GetY()}, nil
	//возвращает response  с атрибутом result
}
