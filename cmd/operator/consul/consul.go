package consul

import (
	"github.com/urfave/cli"
	operator "github.com/zdq0394/consul-operator/operator/consul"
)

// Flags of sub command `consul`
var Flags []cli.Flag

func init() {
	Flags = []cli.Flag{
		cli.BoolFlag{
			EnvVar: "DEVELOP",
			Name:   "develop",
			Usage:  "start the operator in develop mode",
		},
		cli.StringFlag{
			EnvVar: "KUBECONFIG",
			Name:   "kubeconfig",
			Usage:  "kubeconfig of the kubernetes cluster",
		},
		cli.StringFlag{
			EnvVar: "ClusterDomain",
			Name:   "clusterdomain",
			Value:  "cluster.local",
			Usage:  "Kubernetes cluster domain: e.g. cluster.local",
		},
		cli.IntFlag{
			EnvVar: "ConcurrentWorkers",
			Name:   "concurrentworkers",
			Value:  3,
			Usage:  "Concurrent goroutines to process crd management",
		},
	}
}

// Action of sub command `consul`
func Action(ctx *cli.Context) {
	conf := operator.Config{}
	conf.Development = ctx.Bool("develop")
	conf.Kubeconfig = ctx.String("kubeconfig")
	conf.ClusterDomain = ctx.String("clusterdomain")
	conf.ConcurrentWorkers = ctx.Int("concurrentworkers")
	operator.Start(&conf)
}

// Command Consul Sub Command
func Command() cli.Command {
	return cli.Command{
		Name:    "consul",
		Aliases: []string{"c"},
		Usage:   "start consul operator",
		Flags:   Flags,
		Action:  Action,
	}
}
