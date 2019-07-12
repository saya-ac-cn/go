package main

import (
	"go-start/dao"
	"go-start/seelog/util"
	"net/rpc"
)

func main() {
	log := util.GetLoggerInstance().Logger
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:2181")
	if err != nil {
		log.Errorf("dailing error: ", err)
	}

	req := dao.ArithRequest{9, 2}
	var res dao.ArithResponse

	err = client.Call("Arith1.Multiply", req, &res) // 乘法运算
	if err != nil {
		log.Errorf("arith error: ", err)
	}
	log.Debugf("%d * %d = %d\n", req.A, req.B, res.Pro)

	err = client.Call("Arith1.Divide", req, &res)
	if err != nil {
		log.Errorf("arith error: ", err)
	}
	log.Debugf("%d / %d, quo is %d, rem is %d\n", req.A, req.B, res.Quo, res.Rem)
}
