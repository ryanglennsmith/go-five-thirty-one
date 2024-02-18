package cmd

import (
	"fmt"
	"go-five-thirty-one/config"
	googleapi "go-five-thirty-one/internal/google_api"

	"github.com/spf13/cobra"
)

// testDownloadCmd represents the testDownload command
var testDownloadCmd = &cobra.Command{
	Use:   "test-download",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("let's see?")
		doDownload()
	},
}





func doDownload() {
	config := config.GetConfig()
	
	driveService := googleapi.NewDriveService(config.SecretsPath)
	err := driveService.DownloadFile(config.FileId, config.DataFile)
	if err != nil {
		fmt.Println("error downloading file:", err)
		return
	}
	fmt.Printf("file downloaded successfully at %v", config.DataFile)
}

func init() {
	rootCmd.AddCommand(testDownloadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testDownloadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testDownloadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
