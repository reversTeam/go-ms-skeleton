package ping

import (
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/reversTeam/go-ms-skeleton/services/ping/protobuf"
	"github.com/reversTeam/go-ms-tools/services/abs"
	pbAbs "github.com/reversTeam/go-ms-tools/services/abs/protobuf"
	"github.com/reversTeam/go-ms/core"
	"golang.org/x/net/context"
)

// Define the service structure
type Service struct {
	*abs.Service
	pb.UnimplementedPingServer
}

// Instanciate the service without dependency because it's role of ServiceFactory
func NewService(name string, config core.ServiceConfig) *Service {
	s := &Service{
		Service: abs.NewService(name, config),
	}

	return s
}

// This method is required for redister your service on the Http server
func (o *Service) RegisterHttp(gh *core.GoMsHttpServer, endpoint string) error {
	return pb.RegisterPingHandlerFromEndpoint(gh.Ctx, gh.Mux, endpoint, gh.Grpc.Opts)
}

// This method is required for redister your service on the Grpc server
func (o *Service) RegisterGrpc(gs *core.GoMsGrpcServer) {
	pb.RegisterPingServer(gs.Server, o)
}

// Endpoint :
//   - grpc : List
//   - http : Get /ping
func (o *Service) List(ctx context.Context, in *empty.Empty) (*pbAbs.Response, error) {
	return &pbAbs.Response{
		Message: "Ping List",
	}, nil
}

// Endpoint :
//   - grpc : Create
//   - http : POST /ping
func (o *Service) Create(ctx context.Context, in *empty.Empty) (*pbAbs.Response, error) {
	return &pbAbs.Response{
		Message: "Ping Create",
	}, nil
}

// Endpoint :
//   - grpc : Get
//   - http : GET /ping/{id}
func (o *Service) Get(ctx context.Context, in *pbAbs.EntityRequest) (*pbAbs.Response, error) {
	return &pbAbs.Response{
		Message: "Ping View",
	}, nil
}

// Endpoint :
//   - grpc : Update
//   - http : PATCH /ping/{id}
func (o *Service) Update(ctx context.Context, in *pbAbs.EntityRequest) (*pbAbs.Response, error) {
	return &pbAbs.Response{
		Message: "Ping Update",
	}, nil
}

// Endpoint :
//   - grpc : Delete
//   - http : PATCH /ping/{id}
func (o *Service) Delete(ctx context.Context, in *pbAbs.EntityRequest) (*pbAbs.Response, error) {
	return &pbAbs.Response{
		Message: "Ping Delete",
	}, nil
}
