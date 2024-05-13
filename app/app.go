package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Build() *cli.App {
	app := cli.NewApp()

	app.Name = "App de Linha de Comando"
	app.Usage = "Busca ip de servidores e nameservers por domínio"

	Flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "github.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs de um servidor",
			Flags:  Flags,
			Action: buscarIP,
		},
		{
			Name:   "ns",
			Usage:  "Busca os NameServers de um servidor",
			Flags:  Flags,
			Action: buscarNS,
		},
	}

	return app
}

func buscarIP(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarNS(c *cli.Context) {
	host := c.String("host")

	nameServers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ns := range nameServers {
		fmt.Println(ns.Host)
	}
}
