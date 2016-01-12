// Entry point for Twistlock authorization plugin
package main

import (
	"authz/authz"
	"authz/core"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/docker/docker/vendor/src/github.com/Sirupsen/logrus"
	"os"
)

const (
	debugFlag      = "debug"
	handlerFlag    = "authz-handler"
	auditorFlag    = "auditor"
	policyFileFlag = "policy-file"
)

const (
	handlerBasic = "basic"
)

const (
	auditorBasic = "basic"
)

func main() {

	app := cli.NewApp()
	app.Name = "twistlock-authz"
	app.Usage = "Authorization plugin for docker"
	app.Version = "1.0"

	app.Action = func(c *cli.Context) {

		initLogger(c.GlobalBool(debugFlag))

		var auditor core.Auditor
		var authZHandler core.Authorizer

		switch c.GlobalString(handlerFlag) {
		case handlerBasic:
			authZHandler = authz.NewBasicAuthZHandler(&authz.BasicAuthorizerSettings{PolicyPath: c.GlobalString(policyFileFlag)})
		default:
			panic(fmt.Sprintf("Unkwon authz hander %q", c.GlobalString(handlerFlag)))
		}

		switch c.GlobalString(auditorFlag) {
		case auditorBasic:
			auditor = authz.NewBasicAuditor()
		default:
			panic(fmt.Sprintf("Unkwon authz hander %q", c.GlobalString(handlerFlag)))
		}

		srv := core.NewAuthZSrv(authZHandler, auditor)
		err := srv.Start()

		if err != nil {
			panic(err)
		}
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:   debugFlag,
			Usage:  "Enable debug mode",
			EnvVar: "DEBUG",
		},

		cli.StringFlag{
			Name:   handlerFlag,
			Value:  handlerBasic,
			EnvVar: "AUTHZ-HANDLER",
			Usage:  "Defines the authz handler type",
		},

		cli.StringFlag{
			Name:   policyFileFlag,
			Value:  "/var/lib/authz_broker/policy.json",
			EnvVar: "AUTHZ-POLICY-FILE",
			Usage:  "Defines the authz policy file for basic handler",
		},

		cli.StringFlag{
			Name:   auditorFlag,
			Value:  auditorBasic,
			EnvVar: "AUTHZ-AUDITOR",
			Usage:  "Defines the authz auditor type",
		},
	}

	app.Run(os.Args)
}

// initLogger initialize the logger based on the log level
func initLogger(debug bool) {

	logrus.SetFormatter(&logrus.TextFormatter{})
	// Output to stderr instead of stdout, could also be a file.
	logrus.SetOutput(os.Stdout)
	// Only log the warning severity or above.
	logrus.SetLevel(logrus.DebugLevel)
	if debug {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
}
