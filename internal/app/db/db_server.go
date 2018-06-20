package db

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"hello/internal/pkg/pb/pbgame"
	"log"
	"net"
	"time"
)

type DbServer struct {
	cfg *Configuration
	db  *sql.DB
}

func NewDbServer(path string) (*DbServer, error) {
	cfg, err := NewConfiguration(path)
	if err != nil {
		return nil, err
	}

	srv := &DbServer{
		cfg: cfg,
	}

	strConn := fmt.Sprintf("%s:%s@%s(%s)/%s", cfg.Db.User, cfg.Db.Password, cfg.Db.Protocol, cfg.Db.Address, cfg.Db.DbName)

	srv.db, err = sql.Open(cfg.Db.Driver, strConn)
	srv.db.SetMaxOpenConns(1)

	if err != nil {
		return nil, err
	}

	return srv, nil
}

func (db *DbServer) Run() {
	log.Printf("configuration %v", db.cfg)
	lis, err := net.Listen("tcp", db.cfg.Addr)
	if err != nil {
		log.Fatalf("failed to listen at %s. %v", db.cfg.Addr, err)
	}
	grpcServer := grpc.NewServer()
	pbgame.RegisterDBServer(grpcServer, db)
	grpcServer.Serve(lis)
}

func (srv *DbServer) LoadPlayer(context context.Context, req *pbgame.LoadRequest) (*pbgame.LoadResponse, error) {
	rsp := &pbgame.LoadResponse{}
	query, err := srv.db.Prepare("select `id`,`name`,`money` from role where `account`= ?")
	if err != nil {
		rsp.Result = pbgame.ErrorCode_MYSQL_ERROR
		return rsp, err
	}
	defer query.Close()

	err = query.QueryRow(req.Account).Scan(&rsp.Uid, &rsp.Name, &rsp.Money)
	if err != nil {
		rsp.Result = pbgame.ErrorCode_ACCOUNT_NOT_EXISTS
		return rsp, nil
	}

	return rsp, nil
}

func (srv *DbServer) SavePlayer(context context.Context, req *pbgame.SaveRequest) (*pbgame.SaveResponse, error) {
	start := time.Now().UnixNano()
	rsp := &pbgame.SaveResponse{}
	update, err := srv.db.Prepare("UPDATE `role` set `money` = ? where `id`= ?")
	if err != nil {
		rsp.Result = pbgame.ErrorCode_MYSQL_ERROR
		return rsp, err
	}
	defer update.Close()
	_, err = update.Exec(req.Money, req.Uid)
	if err != nil {
		rsp.Result = pbgame.ErrorCode_MYSQL_ERROR
		return rsp, err
	}
	rsp.Result = pbgame.ErrorCode_SUCCESS
	end := time.Now().UnixNano()
	log.Printf("do_save_svc %d %d %d %d", req.Uid, start/1e6, end/1e6, (end-start)/1e6)
	return rsp, nil
}

func (srv *DbServer) CreatePlayer(context context.Context, req *pbgame.CreatePlayerRequest) (*pbgame.CreatePlayerResponse, error) {
	rsp := &pbgame.CreatePlayerResponse{}
	tx, err := srv.db.Begin()
	defer tx.Commit()
	if err != nil {
		rsp.Result = pbgame.ErrorCode_MYSQL_ERROR
		return rsp, err
	}

	query, err := tx.Prepare("select id from `role` where account=? ")
	if err != nil {
		rsp.Result = pbgame.ErrorCode_MYSQL_ERROR
		return rsp, err
	}

	defer query.Close()

	var id uint32
	err = query.QueryRow(req.Account).Scan(&id)
	//account exists
	if err == nil {
		rsp.Result = pbgame.ErrorCode_ACCOUNT_EXISTS
		return rsp, err
	}

	//do insert
	ins, err := tx.Prepare("insert into role(`account`,`name`,`money`) values (?,?,?)")
	if err != nil {
		rsp.Result = pbgame.ErrorCode_MYSQL_ERROR
		return rsp, err
	}
	defer ins.Close()
	result, err := ins.Exec(req.Account, req.Name, req.Money)

	if err != nil {
		rsp.Result = pbgame.ErrorCode_MYSQL_ERROR
		return rsp, err
	}

	rsp.Result = pbgame.ErrorCode_SUCCESS
	uid, err := result.LastInsertId()
	if err != nil {
		rsp.Result = pbgame.ErrorCode_MYSQL_ERROR
		return rsp, err
	}

	rsp.Result = pbgame.ErrorCode_SUCCESS
	rsp.Uid = uint32(uid)
	return rsp, nil
}
