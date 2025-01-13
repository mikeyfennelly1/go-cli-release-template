package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

const MPOD_VERSION = "dev"

func main() {
	var rootCmd = &cobra.Command{
		Use:   "mpod",
		Short: "A CLI tool for creating and managing Linux containers",
	}

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version of mpod",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("mpod %s\n", MPOD_VERSION)
		},
	}

	//var psCmd = &cobra.Command{
	//	Use:   "ps",
	//	Short: "list running containers",
	//	Run: func(cmd *cobra.Command, args []string) {
	//		ps()
	//	},
	//}

	var empty = &cobra.Command{
		Use:   "empty",
		Short: "create an empty container",
		Run: func(cmd *cobra.Command, args []string) {
			StartEmpty()
		},
	}

	rootCmd.AddCommand(versionCmd)
	//rootCmd.AddCommand(psCmd)
	rootCmd.AddCommand(empty)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
