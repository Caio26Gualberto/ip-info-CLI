package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "CLI Ip Finder"
	app.Usage = "Busca IPs e nomes de servidores públicos na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "github.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs públicos na internet",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca o nome dos servidores na internet",
			Flags:  flags,
			Action: searchServers,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, errors := net.LookupIP(host)

	if errors != nil {
		log.Fatal(errors)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServers(c *cli.Context) {
	host := c.String("host")

	servers, errors := net.LookupNS(host)

	if errors != nil {
		log.Fatal(errors)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
