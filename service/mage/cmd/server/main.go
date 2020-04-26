package main

import (
	"fmt"
	"os"

	"gitlab.com/alienspaces/holyragingmages/common/config"
	"gitlab.com/alienspaces/holyragingmages/common/log"
	"gitlab.com/alienspaces/holyragingmages/common/prepare"
	"gitlab.com/alienspaces/holyragingmages/common/server"
	"gitlab.com/alienspaces/holyragingmages/common/store"
	"gitlab.com/alienspaces/holyragingmages/service/mage/internal/server/runner"
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

	p, err := prepare.NewPrepare(l)
	if err != nil {
		fmt.Printf("Failed new prepare >%v<", err)
		os.Exit(0)
	}

	r := runner.NewRunner()

	svr, err := server.NewServer(c, l, s, p, r)
	if err != nil {
		fmt.Printf("Failed new service >%v<", err)
		os.Exit(0)
	}

	args := make(map[string]interface{})

	err = svr.Run(args)
	if err != nil {
		fmt.Printf("Failed service run >%v<", err)
		os.Exit(0)
	}

	os.Exit(1)
}
