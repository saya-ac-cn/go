package main

import (
	"errors"
	"go-start/dao"
	"go-start/seelog/util"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

// https://studygolang.com/articles/14336

// 算数运算结构体
type Arith1 struct {
}

// 乘法运算方法
func (this *Arith1) Multiply(req dao.ArithRequest, res *dao.ArithResponse) error {
	res.Pro = req.A * req.B
	return nil
}

// 除法运算方法
func (this *Arith1) Divide(req dao.ArithRequest, res *dao.ArithResponse) error {
	if req.B == 0 {
		return errors.New("divide by zero")
	}
	res.Quo = req.A / req.B
	res.Rem = req.A % req.B
	return nil
}

func main() {
	log := util.GetLoggerInstance().Logger
	rpc.RegisterName("Arith1", new(Arith1)) // 注册rpc服务
	rpc.HandleHTTP()                        // 采用http协议作为rpc载体

	listener, err := net.Listen("tcp", "127.0.0.1:2181")
	if err != nil {
		log.Warnf("ListenTCP error: ", err)
	}

	log.Debugf("%d start connection", os.Stdout)
	http.Serve(listener, nil)
}
