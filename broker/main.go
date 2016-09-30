// broker consists of the entry point for the twistlock authz broker
package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/Sirupsen/logrus"
	"github.com/twistlock/authz/authz"
	"github.com/twistlock/authz/core"
	"os"
)

const (
	debugFlag       = "debug"
	authorizerFlag  = "authz-handler"
	auditorFlag     = "auditor"
	auditorHookFlag = "auditor-hook"
	policyFileFlag  = "policy-file"
)

const (
	authorizerBasic = "basic"
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

		switch c.GlobalString(authorizerFlag) {
		case authorizerBasic:
			authZHandler = authz.NewBasicAuthZAuthorizer(&authz.BasicAuthorizerSettings{PolicyPath: c.GlobalString(policyFileFlag)})
		default:
			panic(fmt.Sprintf("Unknown authz handler %q", c.GlobalString(authorizerFlag)))
		}

		switch c.GlobalString(auditorFlag) {
		case auditorBasic:
			auditor = authz.NewBasicAuditor(&authz.BasicAuditorSettings{LogHook: c.GlobalString(auditorHookFlag)})
		default:
			panic(fmt.Sprintf("Unknown authz handler %q", c.GlobalString(authorizerFlag)))
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
			Name:   authorizerFlag,
			Value:  authorizerBasic,
			EnvVar: "AUTHORIZER",
			Usage:  "Defines the authz handler type",
		},

		cli.StringFlag{
			Name:   policyFileFlag,
			Value:  "/var/lib/authz-broker/policy.json",
			EnvVar: "AUTHZ-POLICY-FILE",
			Usage:  "Defines the authz policy file for basic handler",
		},

		cli.StringFlag{
			Name:   auditorFlag,
			Value:  auditorBasic,
			EnvVar: "AUDITOR",
			Usage:  "Defines the authz auditor type",
		},
		cli.StringFlag{
			Name:   auditorHookFlag,
			Value:  authz.AuditHookStdout,
			EnvVar: "AUDITOR-HOOK",
			Usage:  "Defines the authz auditor hook type (log engine)",
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
