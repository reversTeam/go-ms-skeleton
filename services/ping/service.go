package ping

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/reversTeam/go-ms/core"
	pb "github.com/reversTeam/go-ms-skeleton/services/ping/protobuf"
	"github.com/reversTeam/go-ms/services/goms"
	ms "github.com/reversTeam/go-ms/services/goms/protobuf"
	"golang.org/x/net/context"
)

// Define the service structure
type PingService struct {
	*goms.GoMsService
}

// Instanciate the service without dependency because it's role of ServiceFactory
func NewService(name string) *PingService {
	s := &PingService{
		GoMsService: goms.NewService(name),
	}

	return s
}

// This method is required for redister your service on the Http server
func (o *PingService) RegisterHttp(gh *core.GoMsHttpServer, endpoint string) error {
	return pb.RegisterPingHandlerFromEndpoint(gh.Ctx, gh.Mux, endpoint, gh.Grpc.Opts)
}

// This method is required for redister your service on the Grpc server
func (o *PingService) RegisterGrpc(gs *core.GoMsGrpcServer) {
	pb.RegisterPingServer(gs.Server, o)
}

// Endpoint :
//  - grpc : List
//  - http : Get /ping
func (o *PingService) List(ctx context.Context, in *empty.Empty) (*ms.GoMsResponse, error) {
	return &ms.GoMsResponse{
		Message: "Ping List",
	}, nil
}

// Endpoint :
//  - grpc : Create
//  - http : POST /ping
func (o *PingService) Create(ctx context.Context, in *empty.Empty) (*ms.GoMsResponse, error) {
	return &ms.GoMsResponse{
		Message: "Ping Create",
	}, nil
}

// Endpoint :
//  - grpc : Get
//  - http : GET /ping/{id}
func (o *PingService) Get(ctx context.Context, in *ms.GoMsEntityRequest) (*ms.GoMsResponse, error) {
	return &ms.GoMsResponse{
		Message: "Ping View",
	}, nil
}

// Endpoint :
//  - grpc : Update
//  - http : PATCH /ping/{id}
func (o *PingService) Update(ctx context.Context, in *ms.GoMsEntityRequest) (*ms.GoMsResponse, error) {
	return &ms.GoMsResponse{
		Message: "Ping Update",
	}, nil
}

// Endpoint :
//  - grpc : Delete
//  - http : PATCH /ping/{id}
func (o *PingService) Delete(ctx context.Context, in *ms.GoMsEntityRequest) (*ms.GoMsResponse, error) {
	return &ms.GoMsResponse{
		Message: "Ping Delete",
	}, nil
}
