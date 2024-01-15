/*
Copyright © 2023 jaronnie <jaron@jaronnie.com>

*/

package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jaronnie/ragingcd/core/server/middlewares"

	"github.com/jaronnie/ragingcd/core/server/engine"

	"github.com/spf13/cobra"

	"github.com/jaronnie/ragingcd/core/server/initialize"
	"github.com/jaronnie/ragingcd/core/server/routers"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "core server",
	Long:  `core server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		middlewares.Use(engine.ServerEngine)
		routers.Router(engine.ServerEngine)
		base := fmt.Sprintf("%s:%s", "0.0.0.0", "8081")

		initialize.Initialize()
		go func() {
			if err := engine.ServerEngine.Run(base); err != nil {
				panic(err)
			}

		}()

		// Wait for interrupt signal to gracefully shutdown the server with
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		select {
		case <-ctx.Done():
			return nil
		default:
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
