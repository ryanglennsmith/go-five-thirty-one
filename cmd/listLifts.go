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
		listLifts(cmd, args)
	},
}

var liftsToDisplay []string

func init() {
	rootCmd.AddCommand(listLiftsCmd)

	// Add flags for specific lifts
	listLiftsCmd.Flags().StringSliceVarP(&liftsToDisplay, "lifts", "l", nil, "Specify lifts to display (comma-separated)")


	// Example: go run main.go list-lifts -lifts=DL,SQ OR go run main.go ll -l OHP,BP
}

func listLifts(cmd *cobra.Command, args []string) {
	lifts, err := cmd.Flags().GetStringSlice("lifts")
	if err != nil {
		fmt.Println("Error getting lifts:", err)
		return
	}

	if len(lifts) > 0 {
		displaySpecifiedLifts(lifts)
		return
	}

	displayAllLifts()
}

func displayAllLifts() {
	for lift, oneRM := range LiftData {
		wc := calculators.NewWeightCalculator(lift, float64(oneRM))
		displayLift(wc, lift)
	}
}

func displaySpecifiedLifts(lifts []string) {
	for _, lift := range lifts {
		lift = strings.ToUpper(lift)
		oneRM, ok := LiftData[lift]
		if !ok {
			fmt.Printf("That ain't a lift: %s\n", lift)
			continue
		}
		wc := calculators.NewWeightCalculator(lift, float64(oneRM))
		displayLift(wc, lift)
	}
}

func displayLift(wc calculators.WeightCalculator, lift string) {
	
	for _, modifier := range calculators.WeightModifiers {
		workingSetWeight, err := wc.CalculateWeight(modifier)
		if err != nil {
			fmt.Println(err)
		}
		var formattedWeight string = fmt.Sprintf("%s - %s: %s\n", lift, modifier.Set, formatWeight(workingSetWeight))
		fmt.Print(formattedWeight)
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