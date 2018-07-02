package dbredis

import (
	"context"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"google.golang.org/grpc"
	"log"
	"net"
	"starter-kit/internal/pkg/pb/pbgame"
	"starter-kit/internal/pkg/util"
)

type DbServer struct {
	cfg       *Configuration
	redisPool *redis.Pool
}

func NewDbServer(path string) (*DbServer, error) {
	cfg, err := NewConfiguration(path)
	if err != nil {
		return nil, err
	}

	srv := &DbServer{
		cfg: cfg,
	}

	host := fmt.Sprintf("%s:%d", cfg.Db.Host, cfg.Db.Port)

	srv.redisPool = &redis.Pool{
		MaxIdle:   10,
		MaxActive: 100,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host, redis.DialPassword(cfg.Db.Password), redis.DialDatabase(cfg.Db.Db))
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
	return srv, nil
}

func (db *DbServer) Run() error {
	log.Printf("configuration %v", db.cfg)
	lis, err := net.Listen("tcp", db.cfg.Addr)
	if err != nil {
		log.Printf("failed to listen at %s. %v", db.cfg.Addr, err)
	}
	grpcServer := grpc.NewServer()
	pbgame.RegisterDBServer(grpcServer, db)
	return grpcServer.Serve(lis)
}

type AppError struct {
	Code pbgame.ErrorCode
	Tip  string
}

func (err AppError) Error() string {
	return fmt.Sprintf("code=%d, %s", err.Code, err.Tip)
}

func (srv *DbServer) LoadPlayer(context context.Context, req *pbgame.LoadRequest) (*pbgame.LoadResponse, error) {
	p := Player{
		Account: req.Account,
	}

	conn := srv.redisPool.Get()
	defer util.Close(conn)

	err := p.Load(conn)

	if err != nil {
		return &pbgame.LoadResponse{Result: pbgame.ErrorCode_ACCOUNT_NOT_EXISTS}, nil
	}

	rsp := &pbgame.LoadResponse{
		Uid:   p.Id,
		Name:  p.Name,
		Money: int32(p.Money),
	}

	return rsp, nil
}

func (srv *DbServer) SavePlayer(context context.Context, req *pbgame.SaveRequest) (*pbgame.SaveResponse, error) {
	p := Player{
		Account: req.Account,
		Name:    req.Name,
		Id:      req.Uid,
		Money:   int64(req.Money),
	}

	conn := srv.redisPool.Get()
	defer util.Close(conn)

	err := p.Save(conn)

	if err != nil {
		return nil, err
	}

	rsp := &pbgame.SaveResponse{Result: pbgame.ErrorCode_SUCCESS}
	return rsp, nil
}

func (srv *DbServer) CreatePlayer(context context.Context, req *pbgame.CreatePlayerRequest) (*pbgame.CreatePlayerResponse, error) {
	rsp, err := srv.doCreatePlayer(context, req)

	if appErr, ok := err.(AppError); ok {
		rsp = &pbgame.CreatePlayerResponse{Result: appErr.Code}
	}

	log.Printf("CreatePlayer return %v %v", rsp, err)
	return rsp, err
}
func (srv *DbServer) doCreatePlayer(ctx context.Context, req *pbgame.CreatePlayerRequest) (*pbgame.CreatePlayerResponse, error) {
	conn := srv.redisPool.Get()
	defer util.Close(conn)

	ok, err := redis.Bool(conn.Do("HEXISTS", PlayerKey, req.Account))

	if err != nil {
		return nil, AppError{Code: pbgame.ErrorCode_MYSQL_ERROR}
	}

	if ok {
		return nil, AppError{Code: pbgame.ErrorCode_ACCOUNT_EXISTS}
	}

	id, err := redis.Uint64(conn.Do("INCR", AutoAccountId))

	if err != nil {
		return nil, AppError{Code: pbgame.ErrorCode_ACCOUNT_EXISTS}
	}

	p := &Player{
		Account: req.Account,
		Id:      uint32(id),
		Name:    req.Name,
		Money:   int64(req.Money),
	}

	err = p.Save(conn)
	if err != nil {
		return nil, AppError{Code: pbgame.ErrorCode_MYSQL_ERROR}
	}

	rsp := &pbgame.CreatePlayerResponse{
		Result: pbgame.ErrorCode_SUCCESS,
		Uid:    p.Id,
	}

	return rsp, nil
}
