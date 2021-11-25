package main

import (
	"flag"
	"fmt"
	"go-crontab/master"
	"runtime"
)

var (
	fileName string
)

func initArgs()  {
	flag.StringVar(&fileName, "config", "./master.json", "指定master.json")
	flag.Parse()
}

func initEnv()  {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main()  {
	var (
		err error
	)
	// 初始化
	initArgs()
	initEnv()
	if err = master.InitConfig(fileName);err != nil {
		goto ERR
	}
	// 任务管理器
	if err = master.InitJobMgr();err != nil {
		goto ERR
	}

	// 启动HTTP服务
	if err = master.InitApiServer(); err != nil {
		goto ERR
	}
	ERR:
		fmt.Println(err)
}