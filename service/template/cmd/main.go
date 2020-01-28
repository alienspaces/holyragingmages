package main

import (
	"fmt"
	"os"

	"gitlab.com/alienspaces/holyragingmages/common/config"
	"gitlab.com/alienspaces/holyragingmages/common/database"
	"gitlab.com/alienspaces/holyragingmages/common/logger"
	"gitlab.com/alienspaces/holyragingmages/common/service"
	"gitlab.com/alienspaces/holyragingmages/service/template/internal/runner"
)

func main() {

	c, err := config.NewConfig(nil, false)
	if err != nil {
		fmt.Printf("Failed new config >%v<", err)
		os.Exit(0)
	}

	l, err := logger.NewLogger(c)
	if err != nil {
		fmt.Printf("Failed new logger >%v<", err)
		os.Exit(0)
	}

	d, err := database.NewDatabase(c, l)
	if err != nil {
		fmt.Printf("Failed new database >%v<", err)
		os.Exit(0)
	}

	r := runner.NewRunner()

	s, err := service.NewService(c, l, d, r)
	if err != nil {
		fmt.Printf("Failed new service >%v<", err)
		os.Exit(0)
	}

	args := make(map[string]interface{})

	s.Run(args)

	os.Exit(1)
}
