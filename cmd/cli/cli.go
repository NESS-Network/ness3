package main

import (
	"fmt"
	"os"

	"github.com/MDLlife/MDL/src/api/cli"
)

func main() {
	cfg, err := cli.LoadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	app := cli.NewApp(cfg)

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
