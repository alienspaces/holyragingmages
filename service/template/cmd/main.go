package main

import (
	"fmt"
	"os"

	"gitlab.com/alienspaces/holyragingmages/common/database"
	"gitlab.com/alienspaces/holyragingmages/common/env"
	"gitlab.com/alienspaces/holyragingmages/common/logger"
	"gitlab.com/alienspaces/holyragingmages/common/service"

	"gitlab.com/alienspaces/holyragingmages/service/template/internal/template"
)

func main() {

	e, err := env.NewEnv(nil, true)
	if err != nil {
		fmt.Printf("Failed new env >%v<", err)
		os.Exit(0)
	}

	l, err := logger.NewLogger(e)
	if err != nil {
		fmt.Printf("Failed new logger >%v<", err)
		os.Exit(0)
	}

	d, err := database.NewDatabase(e, l)
	if err != nil {
		fmt.Printf("Failed new database >%v<", err)
		os.Exit(0)
	}

	r := template.Runner{}

	s, err := service.NewService(e, l, d, &r)
	if err != nil {
		fmt.Printf("Failed new service >%v<", err)
		os.Exit(0)
	}

	s.Run()

	os.Exit(1)
}
