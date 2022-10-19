package main

import (
	"flag"
	"fmt"
	"log"
	"mongodb_query/config"
	"mongodb_query/db"
	"mongodb_query/query"
	"os"
	"os/signal"
	"syscall"
)

var configFlag = flag.String("config", "./config.toml", "configuration toml file path")

type App struct {
	stop     chan interface{}
	database *db.Database
}

func New() (*App, error) {
	app := &App{
		stop: make(chan interface{}),
	}

	config := config.NewConfig(*configFlag)
	database, err := db.NewDatabase(config)
	if err != nil {
		return nil, err
	}
	app.database = database

	if tx, err := query.FindTxByHash(database.TxCollection, "0x075164408b59135a8efd2dc840147d397007552b92e14a2ca79e60d8b0d17f98"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(tx)
	}

	if txs, err := query.GetTxsByBlockNumberGT(database.TxCollection, 7923820); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(txs)
	}

	go func() {
		sigc := make(chan os.Signal, 1)
		signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(sigc)
		defer close(sigc)
		<-sigc

		go app.Terminate()
	}()

	return app, nil
}

func (p *App) Wait() {
	<-p.stop
}

func (p *App) Terminate() {
	defer close(p.stop)

	if err := p.database.Disconnect(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("몽고DB 연결을 종료했습니다!")
}

func main() {
	if n, err := New(); err != nil {
		panic(err)
	} else {
		n.Wait()
	}
}
