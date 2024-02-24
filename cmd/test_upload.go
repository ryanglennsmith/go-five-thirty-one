package cmd

import (
	"fmt"
	"go-five-thirty-one/config"
	googleapi "go-five-thirty-one/internal/google_api"
	"go-five-thirty-one/internal/util"

	"github.com/spf13/cobra"
)

// testUploadCmd represents the testUpload command
var testUploadCmd = &cobra.Command{
	Use:   "test-upload",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("testUpload called")
		doUpload()
	},
}

func doUpload() {
	config := config.GetConfig()
	done := util.StartLoadingIndicator()
	driveService := googleapi.NewDriveService(config.SecretsPath)
	err := driveService.UpdateFile(config.FileId, config.DataFile)
	if err != nil {
		fmt.Println("error uploading file:", err)
		return
	}
	done <- true
	fmt.Println("uploaded file")
}
func init() {
	rootCmd.AddCommand(testUploadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testUploadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testUploadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
