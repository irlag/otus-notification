package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"otus-notification/app"
	"otus-notification/app/config"
)

var Version = ""

var consoleCmd = &cobra.Command{
	Use:     "otus-notification",
	Short:   "Education otus notification app",
	Long:    ``,
	Version: Version,
}

var container *app.Container

func init() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	container = app.NewContainer(cfg)

	consoleCmd.AddCommand(apiServer)
	consoleCmd.AddCommand(version)
	consoleCmd.AddCommand(comsumeNotification)
}

func Execute() {
	if err := consoleCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
