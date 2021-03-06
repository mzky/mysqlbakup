package module

import (
	"os"
	"os/signal"

	"github.com/mzky/mysqlbakup/common/console"
	"github.com/mzky/mysqlbakup/common/log"
)

var (
	OnDestroy func()
	CloseTag  = make(chan int32, 1)
)

func Run(mods ...Module) {
	// logger
	//if conf.LogLevel != "" {
	//	logger, err := log.New(conf.ServerName, conf.LogLevel, conf.LogPath, conf.LogFlag)
	//	if err != nil {
	//		panic(err)
	//	}
	//	log.Export(logger)
	//	defer logger.Close()
	//}

	log.Release("common %v starting up", 1.0)

	// module
	for i := 0; i < len(mods); i++ {
		Register(mods[i])
	}
	Init()
	// console
	console.Init(CloseTag)

	// close
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, os.Kill)
		sig1 := <-c
		log.Release("Leaf closing down (signal: %v)", sig1)
		var a = int32(1)
		console.CloseTag <- a

	}()

	sig := <-console.CloseTag
	log.Release("Leaf closing down1 (signal: %v)", sig)

	if OnDestroy != nil {
		OnDestroy()
	}
	console.Destroy()
	Destroy()

}
