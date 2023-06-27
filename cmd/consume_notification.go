package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"otus-notification/app/models"
)

var comsumeNotification = &cobra.Command{
	Use:   "consume-notification",
	Short: "Run notification consumer",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		go container.RecipeNotificationReader.Read(func(recipe models.RecipeNotification) error {
			return container.Processors.EventProcessor.HandleRecipeNotification(recipe)
		})

		exit := make(chan os.Signal, 1)
		signal.Notify(exit, syscall.SIGTERM, syscall.SIGINT)

		<-exit

		fmt.Println("Closing Kafka connections ...")
	},
}
