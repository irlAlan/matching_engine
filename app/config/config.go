package config

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Network struct {
	Network string `toml:"network"`
	Ip_addr string `toml:"ip_addr"`
	Port    string `toml:"port"`
}

type Config struct {
	Connection Network `toml:"connection"`
}

func (cc *Config) ReadConfig() {
	filepath := "config.toml"
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	// var cc Config
	doc, err := toml.Decode(string(data), &cc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Undecoded keys: %q\n", doc.Undecoded())

	if !doc.IsDefined("connection", "network") {
		fmt.Printf("conf.network is not defined, using default: tcp\n")
		cc.Connection.Network = "tcp"
	}
	if !doc.IsDefined("connection", "ip_addr") {
		fmt.Printf("connection.ip_addr is not defined, using defualt: localhost\n")
		cc.Connection.Ip_addr = "localhost"
	}
	if !doc.IsDefined("connection", "port") {
		fmt.Printf("connection.port is not defined, using defualt: 50051\n")
		cc.Connection.Port = ":50051"
	}
	// return cc
}
