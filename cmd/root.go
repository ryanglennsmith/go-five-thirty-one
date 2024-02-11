/*
Copyright © 2024 RYAN GLENN SMITH

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-five-thirty-one",
	Short: "a 5/3/1 calculator and tracker and some other stuff",
	Long: `	TODO: calculator for 5/3/1
	TODO: enter/update weights
	TODO: persist data (csv?)
	TODO: calculator for 1RM
	TODO: calculator for plate math
	TODO: setup for plates and bars
	TODO: wire with google api for sheets
	TODO: CRUD tracker for workouts
	TODO: daily cron service to sync with sheets
	TODO: macro nutrition stuff
	TODO: play with ANSI color stuff
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-five-thirty-one.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


