package cmd

import (
	"fmt"
	"go-five-thirty-one/internal/csv"
	"strings"

	"github.com/spf13/cobra"
)

// this file path will become dynamic with cycle changes
const path = "./.csv/data.csv"



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
	userData, err := csv.ReadData(path)
	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}
	
	lifts, err := cmd.Flags().GetStringSlice("lifts")
	if err != nil {
		fmt.Println("Error getting lifts:", err)
		return
	}

	if len(lifts) > 0 {
		displaySpecifiedLifts(lifts, userData)
		return
	}

	displayAllLifts(userData)
}


func displayAllLifts(data []csv.LiftData) {
	for _, obj := range data {
		displayLift(obj)
	}
}

func displaySpecifiedLifts(lifts []string, data []csv.LiftData) {
	for _, lift := range lifts {
		lift = strings.ToUpper(lift)
		
		var slicedData []csv.LiftData
		for _, obj := range data {
			if obj.Lift == lift {
				slicedData = append(slicedData, obj)
			}
		}
		displayAllLifts(slicedData)
		
	}
}


func displayLift(item csv.LiftData) {
	// for padding to display evenly
	liftStrLen := 3
	weightStrLen := 5
	formattedLift := fmt.Sprintf("%-*s", liftStrLen, item.Lift)
	formattedWeight := fmt.Sprintf("%-*s", weightStrLen, formatWeight(item.Weight))
	fmt.Printf("%s - %s: %s", formattedLift, item.WeekSet, formattedWeight)

if item.Date != "" {
    fmt.Printf(" | %s", item.Date)
}

if item.Comments != "" {
    fmt.Printf(" | %s", item.Comments)
}

fmt.Println()
}

func formatWeight(weight float64) string {
	// Convert to string
	weightStr := fmt.Sprintf("%.2f", weight)

	// Remove trailing zeroes and the decimal point if it's a whole number
	weightStr = strings.TrimRight(weightStr, "0")
	weightStr = strings.TrimRight(weightStr, ".")

	return weightStr
}