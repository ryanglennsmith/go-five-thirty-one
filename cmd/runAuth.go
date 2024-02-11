/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"go-five-thirty-one/internal/googleauth"

	"github.com/spf13/cobra"
)

// runAuthCmd represents the runAuth command
var runAuthCmd = &cobra.Command{
	Use:   "run-auth",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("runAuth called")
		googleauth.RunAuth()
	},
}

func init() {
	rootCmd.AddCommand(runAuthCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runAuthCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runAuthCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
