package main

import (
	"fmt"
	"os"

	"gitlab.com/alienspaces/holyragingmages/server/core/cli"
	"gitlab.com/alienspaces/holyragingmages/server/core/config"
	"gitlab.com/alienspaces/holyragingmages/server/core/log"
	"gitlab.com/alienspaces/holyragingmages/server/core/store"
	"gitlab.com/alienspaces/holyragingmages/server/service/template/internal/cli/runner"
)

func main() {

	c, err := config.NewConfig(nil, false)
	if err != nil {
		fmt.Printf("Failed new config >%v<", err)
		os.Exit(0)
	}

	configVars := []string{
		// general
		"APP_SERVER_ENV",
		"APP_SERVER_PORT",
		// logger
		"APP_SERVER_LOG_LEVEL",
		// database
		"APP_SERVER_DB_HOST",
		"APP_SERVER_DB_PORT",
		"APP_SERVER_DB_NAME",
		"APP_SERVER_DB_USER",
		"APP_SERVER_DB_PASSWORD",
	}
	for _, key := range configVars {
		err := c.Add(key, true)
		if err != nil {
			fmt.Printf("Failed adding config item >%v<", err)
			os.Exit(0)
		}
	}

	l, err := log.NewLogger(c)
	if err != nil {
		fmt.Printf("Failed new logger >%v<", err)
		os.Exit(0)
	}

	s, err := store.NewStore(c, l)
	if err != nil {
		fmt.Printf("Failed new store >%v<", err)
		os.Exit(0)
	}

	r := runner.NewRunner()

	cli, err := cli.NewCLI(c, l, s, r)
	if err != nil {
		fmt.Printf("Failed new cli >%v<", err)
		os.Exit(0)
	}

	args := make(map[string]interface{})

	err = cli.Run(args)
	if err != nil {
		fmt.Printf("Failed cli run >%v<", err)
		os.Exit(0)
	}

	os.Exit(1)
}