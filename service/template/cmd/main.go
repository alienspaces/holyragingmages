package main

import (
	"fmt"
	"os"

	"gitlab.com/alienspaces/holyragingmages/common/config"
	"gitlab.com/alienspaces/holyragingmages/common/logger"
	"gitlab.com/alienspaces/holyragingmages/common/service"
	"gitlab.com/alienspaces/holyragingmages/common/store"
	"gitlab.com/alienspaces/holyragingmages/service/template/internal/runner"
)

func main() {

	c, err := config.NewConfig(nil, false)
	if err != nil {
		fmt.Printf("Failed new config >%v<", err)
		os.Exit(0)
	}

	configVars := []string{
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

	l, err := logger.NewLogger(c)
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

	svc, err := service.NewService(c, l, s, r)
	if err != nil {
		fmt.Printf("Failed new service >%v<", err)
		os.Exit(0)
	}

	args := make(map[string]interface{})

	svc.Run(args)

	os.Exit(1)
}
