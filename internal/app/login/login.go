package login

import (
	"database/sql"
	"fmt"
	"hello/internal/pkg/pb/pbgame"
	"log"
	"net"

	"hello/internal/pkg/util"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Login struct {
	cfg *Configuration
	db  *sql.DB
}

func NewLogin(path string) (*Login, error) {
	cfg, err := NewConfiguration(path)
	if err != nil {
		return nil, err
	}

	login := &Login{
		cfg: cfg,
	}

	strConn := fmt.Sprintf("%s:%s@%s(%s)/%s", cfg.Db.User, cfg.Db.Password, cfg.Db.Protocol, cfg.Db.Address, cfg.Db.DbName)

	login.db, err = sql.Open(cfg.Db.Driver, strConn)

	if err != nil {
		return nil, err
	}

	return login, nil
}

func (login *Login) Run() error {
	log.Printf("configuration %v", login.cfg)
	lis, err := net.Listen("tcp", login.cfg.Addr)
	if err != nil {
		log.Fatalf("failed to listen at %s. %v", login.cfg.Addr, err)
	}

	grpcServer := grpc.NewServer()
	pbgame.RegisterLoginServer(grpcServer, login)
	return grpcServer.Serve(lis)
}

func (login *Login) Login(ctx context.Context, req *pbgame.LoginRequest) (*pbgame.LoginResponse, error) {
	//TODO check request

	rsp := &pbgame.LoginResponse{}
	login.doLogin(req, rsp)
	return rsp, nil
}

func (login *Login) doLogin(req *pbgame.LoginRequest, rsp *pbgame.LoginResponse) {
	tx, err := login.db.Begin()
	if err != nil {
		log.Printf("create transaction failed. %v", err)
		rsp.ErrorCode = pbgame.ErrorCode_MYSQL_ERROR
		return
	}

	query, err := tx.Prepare("select password from account where account=?")
	if err != nil {
		log.Printf("prepare sql failed. %v", err)
		rsp.ErrorCode = pbgame.ErrorCode_MYSQL_ERROR
		if err = tx.Rollback(); err != nil {
			log.Fatalf("rollback failed.%v", err)
		}
		return
	}
	defer util.Close(query)

	var password string
	err = query.QueryRow(req.Account).Scan(password)

	//do register
	if err != nil {
		ins, err := tx.Prepare("insert into account(`account`,`password`) values(?,?)")
		if err != nil {
			log.Fatalf("tx prepared faield. %v", err)
			rsp.ErrorCode = pbgame.ErrorCode_MYSQL_ERROR
			if err = tx.Rollback(); err != nil {
				log.Fatalf("rollback failed.%v", err)
			}
			return
		}

		ins.Exec(req.Account, req.Password)

		err = tx.Commit()
		if err != nil {
			log.Printf("tx commit error. %v", err)
			rsp.ErrorCode = pbgame.ErrorCode_MYSQL_ERROR
			return
		}

		rsp.ErrorCode = pbgame.ErrorCode_SUCCESS
		return
	}

	//account has been registered
	if password != req.Password {
		rsp.ErrorCode = pbgame.ErrorCode_LOGIN_PASSWORD_NOT_MATCH
		return
	}

	rsp.ErrorCode = pbgame.ErrorCode_SUCCESS
}
