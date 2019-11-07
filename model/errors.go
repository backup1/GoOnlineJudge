package model

import (
	"GoOnlineJudge/restweb/golog"
	"errors"
	"os"
)

var (
	//数杮库打开错误
	DBErr = errors.New("DB Open Error")

	//函数坂数错误
	ArgsErr = errors.New("Args Error")

	//没有查询到指定数杮
	NotFoundErr = errors.New("Not Found")

	//更新错误
	OpErr = errors.New("Operate Error")

	//ID生戝错误
	IDErr = errors.New("Get ID Error")

	//查询或查询坂数错误
	QueryErr = errors.New("Query Error")

	//密砝加密错误
	EncryptErr = errors.New("Encrypt Error")

	ExistErr = errors.New("Id has existed")
)

var logger *golog.Log

func init() {
	logger = golog.NewLog(os.Stdout, golog.Ldebug|golog.Linfo)
}
