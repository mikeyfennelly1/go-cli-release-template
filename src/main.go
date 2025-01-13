package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "mpod",
		Short: "A CLI tool for creating and managing Linux containers",
	}

	rootCmd.AddCommand(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
