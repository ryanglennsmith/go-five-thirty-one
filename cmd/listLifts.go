package cmd

import (
	"fmt"
	"go-five-thirty-one/internal/calculators"
	"strings"

	"github.com/spf13/cobra"
)

//hardcoded data for now
var LiftData = map[string]int{
	"BP": 100,
	"SQ": 150,
	"DL": 200,
	"OHP": 75,
}

// listLiftsCmd represents the list-lifts command
var listLiftsCmd = &cobra.Command{
	Use:   "list-lifts",
	Aliases: []string{"ll"},
	Short: "lists the lifts and their current weights",
	Long: `TODO: think about weeks?`,
	Run: func(cmd *cobra.Command, args []string) {
		listLifts()
	},
}

func init() {
	rootCmd.AddCommand(listLiftsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listLiftsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listLiftsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")


}

func listLifts() {
	for lift, oneRM := range LiftData {
		wc := calculators.NewWeightCalculator(lift, float64(oneRM))

		for _, modifier := range calculators.WeightModifiers {
			workingSetWeight, err := wc.CalculateWeight(modifier)
			if err != nil {
				fmt.Println(err)
			}
			var formattedWeight string = fmt.Sprintf("%s - %s: %s\n", lift, modifier.Set, formatWeight(workingSetWeight))

			fmt.Print(formattedWeight)}
	}
}

func formatWeight(weight float64) string {
	// Convert to string
	weightStr := fmt.Sprintf("%.2f", weight)

	// Remove trailing zeroes and the decimal point if it's a whole number
	weightStr = strings.TrimRight(weightStr, "0")
	weightStr = strings.TrimRight(weightStr, ".")

	return weightStr
}