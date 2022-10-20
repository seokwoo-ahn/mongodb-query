package main

import (
	"flag"
	"fmt"
	"log"
	"mongodb_query/config"
	"mongodb_query/db"
	"mongodb_query/scanner"
	"os"
	"os/signal"
	"syscall"
)

var configFlag = flag.String("config", "./config.toml", "configuration toml file path")

type App struct {
	stop     chan interface{}
	database *db.Database
	scanner  *scanner.Scanner
}

func New() (*App, error) {
	var err error
	app := &App{
		stop: make(chan interface{}),
	}

	go func() {
		sigc := make(chan os.Signal, 1)
		signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(sigc)
		defer close(sigc)
		<-sigc

		go app.Terminate()
	}()

	config := config.NewConfig(*configFlag)
	if app.database, err = db.NewDatabase(config); err != nil {
		return nil, err
	}
	if app.scanner, err = scanner.NewScanner(app.database); err != nil {
		return nil, err
	}

	return app, nil
}

func (p *App) Wait() {
	<-p.stop
	fmt.Println("앱을 종료합니다")
}

func (p *App) Terminate() {
	defer close(p.stop)
	close(p.scanner.Quit)

	if err := db.Disconnect(p.database); err != nil {
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
