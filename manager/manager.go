package manager

import (
	"errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"hello/pb/pbgame"
	"hello/util"
	"log"
	"net"
)

type Manager struct {
	cfg *Configuration
}

func NewManager(path string) (*Manager, error) {
	cfg, err := NewConfiguration(path)
	if err != nil {
		return nil, err
	}

	mgr := &Manager{
		cfg: cfg,
	}

	return mgr, nil
}

func (mgr *Manager) Run() {
	lis, err := net.Listen("tcp", mgr.cfg.Addr)
	log.Printf("manager listen at %s", mgr.cfg.Addr)
	if err != nil {
		log.Fatalf("failed to listen at %s. %v", mgr.cfg.Addr, err)
	}

	grpcServer := grpc.NewServer()
	pbgame.RegisterManagerServer(grpcServer, mgr)
	grpcServer.Serve(lis)
}

func (mgr *Manager) GetServerList(ctx context.Context, req *pbgame.ServerListReq) (*pbgame.ServerListRsp, error) {
	log.Printf("receive GetServerList request %v", req)
	if req.Token != mgr.cfg.Token {
		return nil, errors.New("invalid token")
	}

	rsp := &pbgame.ServerListRsp{}
	switch req.Server.Type {
	case util.ServerTypeGate:
		mgr.SetServerListForGate(rsp)
		return rsp, nil
	default:
		log.Printf("no server list for %v", req.Server)
	}
	return rsp, nil
}

func (mgr *Manager) SetServerListForGate(rsp *pbgame.ServerListRsp) {
	for _, s := range mgr.cfg.Servers {
		if s.Type == util.ServerTypeGame || s.Type == util.ServerTypeState {
			srv := &pbgame.Server{
				Id:   uint32(s.Id),
				Type: uint32(s.Type),
				Addr: s.Addr,
			}
			rsp.Server = append(rsp.Server, srv)
		}
	}
}
