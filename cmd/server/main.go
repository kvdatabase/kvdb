package main

import (
	"flag"
	"fmt"
	"github.com/pelletier/go-toml"
	"github.com/kvdatabase/kvdb"
	"github.com/kvdatabase/kvdb/cmd"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	// Print banner.
	banner, _ := ioutil.ReadFile("../../resource/banner.txt")
	fmt.Println(string(banner))
}

// The param config means the config file path for kvdb.
// For the default config file, see config.toml.
var config = flag.String("config", "", "the config file for kvdb")

// The param dirPath means the persistent directory for db files and other config.
var dirPath = flag.String("dir_path", "", "the dir path for the database")

func main() {
	flag.Parse()

	// Set the config.
	var cfg kvdb.Config
	if *config == "" {
		log.Println("no config set, using the default config.")
		cfg = kvdb.DefaultConfig()
	} else {
		c, err := newConfigFromFile(*config)
		if err != nil {
			log.Printf("load config err : %+v\n", err)
			return
		}
		cfg = *c
	}

	if *dirPath == "" {
		log.Println("no dir path set, using the os tmp dir.")
	} else {
		cfg.DirPath = *dirPath
	}

	// Listen the server.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill, syscall.SIGHUP,
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	db, err := kvdb.Open(cfg)
	if err != nil {
		log.Printf("create kvdb err: %+v\n", err)
		return
	}

	server := cmd.NewServerUseDbPtr(db)
	grpcServer := cmd.NewGrpcServer(db)
	go server.Listen(cfg.Addr)
	go grpcServer.Listen(cfg.GrpcAddr)
	<-sig
	server.Stop()
	log.Println("kvdb is ready to exit, bye...")
}

func newConfigFromFile(config string) (*kvdb.Config, error) {
	data, err := ioutil.ReadFile(config)
	if err != nil {
		return nil, err
	}

	var cfg = new(kvdb.Config)
	err = toml.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
