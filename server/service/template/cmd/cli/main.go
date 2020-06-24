package main

import (
	"fmt"
	"os"

	"gitlab.com/alienspaces/holyragingmages/server/common/cli"
	"gitlab.com/alienspaces/holyragingmages/server/common/config"
	"gitlab.com/alienspaces/holyragingmages/server/common/log"
	"gitlab.com/alienspaces/holyragingmages/server/common/store"
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
		"APP_ENV",
		"APP_PORT",
		// logger
		"APP_LOG_LEVEL",
		// database
		"APP_DB_HOST",
		"APP_DB_PORT",
		"APP_DB_NAME",
		"APP_DB_USER",
		"APP_DB_PASSWORD",
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
