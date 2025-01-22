package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"

	"github.com/ngtrdai197/go-rabbitmq/config"
	"github.com/ngtrdai197/go-rabbitmq/pkg/helper"
	"github.com/ngtrdai197/go-rabbitmq/pkg/logger"
	"github.com/ngtrdai197/go-rabbitmq/pkg/rabbitmq"
)

func init() {
	rootCmd.AddCommand(userProfileConsumer)
	config.Init()
}

var userProfileConsumer = &cobra.Command{
	Use:   "rabbitmq:consume:user-profile",
	Short: "Serve user profile consumer application",
	Long:  ``,
	Run: func(_ *cobra.Command, _ []string) {
		c, err := rabbitmq.NewConsumer(
			"amqp://guest:guest@localhost:5672/",
			"go.rabbit.exchange",
			"direct",
			"account_profile.created.queue",
			"account_profile.created",
			helper.RandomString("account_profile#"),
		)
		if err != nil {
			panic(err)
		}

		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

		<-signals
		logger.Info(context.Background()).Msg("Shutting down consumer ...")

		logger.Info(context.Background()).Msg("Waiting to consumer can finish in-flight messages ...")
		time.Sleep(10 * time.Second)

		if err := c.Shutdown(); err != nil {
			panic(err)
		}
		logger.Info(context.Background()).Msg("Consumer shutdown successfully")
	},
}
