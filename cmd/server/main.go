package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mikeyobrien/log-aggregator/internal/collector"
	"github.com/spf13/cobra"
)

func main() {
	var path string

	rootCmd := &cobra.Command{
		Use: "log-aggregator",
		Run: func(cmd *cobra.Command, args []string) {
			fc, err := collector.NewFileCollector(path)
			if err != nil {
				log.Fatal(err)
			}
			fc.Start(context.Background())
		},
	}

	// Add flags to root command
	rootCmd.Flags().StringVarP(&path, "path", "p", "default", "path to watch")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
