package dbredis

import (
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"strconv"
)

type Player struct {
	Account string
	Name    string
	Id      uint32
	Money   int64
}

func (p *Player) saveArgs() []interface{} {
	var args []interface{}

	args = append(args, p.redisKey())

	args = append(args, "acc")
	args = append(args, p.Account)

	args = append(args, "id")
	args = append(args, p.Id)

	args = append(args, "name")
	args = append(args, p.Name)

	args = append(args, "money")
	args = append(args, p.Money)
	return args
}

func (p *Player) redisKey() string {
	return fmt.Sprintf("%s-%s", PlayerKey, p.Account)
}

func (p *Player) Save(conn redis.Conn) error {
	args := p.saveArgs()
	_, err := conn.Do("HMSET", args...)
	return err
}

func (p *Player) Load(conn redis.Conn) error {
	args, err := redis.StringMap(conn.Do("HGETALL", p.redisKey()))
	if err != nil {
		return err
	}

	if id, ok := args["id"]; ok {
		intId, err := strconv.Atoi(id)
		if err != nil {
			return err
		}

		p.Id = uint32(intId)
	} else {
		return errors.New(fmt.Sprintf("id not found in map, account=%s", p.Account))
	}

	if name, ok := args["name"]; ok {
		p.Name = name
	} else {
		return errors.New(fmt.Sprintf("name not found in map, account=%s", p.Account))
	}

	if money, ok := args["money"]; ok {
		intMoney, err := strconv.Atoi(money)
		if err != nil {
			return err
		}

		p.Money = int64(intMoney)
	} else {
		return errors.New(fmt.Sprintf("money not found in map, account=%s", p.Account))
	}
	return nil
}
