package main

import (
	"fmt"
	"os"

	goenv "github.com/tangx/goenv"
)

func main() {
	read()
	dump()
}

var (
	cfgfile = "config.yml"
	appname = "AppName"
)

func read() {

	server := &Server{
		Address: "0.0.0.0",
	}

	config := &struct {
		Server *Server
	}{
		Server: server,
	}

	err := goenv.UnmarshalFile(config, appname, cfgfile)
	if err != nil {
		panic(err)
	}

	fmt.Println("addr=", config.Server.Address)
	fmt.Println("port=", config.Server.Port)

	// addr= 192.168.100.100
	// port= 80
}

func dump() {

	server := &Server{
		Address: "192.168.100.100",
	}

	config := &struct {
		Server *Server
	}{
		Server: server,
	}

	err := goenv.CallSetDefaults(config)
	if err != nil {
		panic(err)
	}

	b, err := goenv.Marshal(config, appname)
	if err != nil {
		panic(err)
	}
	_ = os.WriteFile("default.yml", b, os.ModePerm)
}

type Server struct {
	Address string `env:"address,expandenv"`
	Port    int    `env:"port"`
}

// SetDefaults 设置默认值
func (addr *Server) SetDefaults() {
	if addr.Address == "" {
		addr.Address = "localhost"
	}

	if addr.Port == 0 {
		addr.Port = 80
	}
}
