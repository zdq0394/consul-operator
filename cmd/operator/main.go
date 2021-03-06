package main

import (
	"os"

	"github.com/urfave/cli"
	"github.com/zdq0394/consul-operator/cmd/operator/consul"
	"github.com/zdq0394/k8soperator/pkg/log"
)

func main() {
	app := cli.NewApp()
	app.Name = "consul"
	app.Description = "Consul Operator manages the creation/update/deletion of Consul Cluster."
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "DqZhang",
			Email: "zdq123.hn@163.com",
		},
	}
	app.Flags = []cli.Flag{}
	app.Commands = []cli.Command{
		consul.Command(),
	}
	if err := app.Run(os.Args); err != nil {
		log.Errorln(err)
	}
}
