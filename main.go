package main

import (
	"github.com/spf13/cobra"

	"github.com/a-urth/competitive-algo/cmd"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "competitive-algo",
		Short: "Number of algorithms used in competitive programming",
	}

	rootCmd.AddCommand(cmd.SubArrayMaxSumCmd)

	rootCmd.Execute()
}
