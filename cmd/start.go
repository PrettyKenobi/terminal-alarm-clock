/*
Copyright © 2022 Ken Bonnstetter <pretteykenobi@gmail.com>
*/
package cmd

import (
	"time"

	"github.com/spf13/cobra"
)

// Startcmd Represents The Start Command
var startCmd = &cobra.Command{
	Use:   "start name {-t time | -l length}",
	Short: "Start a new timer called name",
	Long: `Start a new timer called name that will finish at time or is to run for length.`,
	Example: `start tea -l 5m
start party -t Nov 5, 2025`,
	ValidArgs: []string{"name"},
	Args: cobra.OnlyValidArgs,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("start called")
		// TODO check if name provided, prompt for one if false
		// TODO check that name is unused, prompt for a different name if false
		// TODO check if duration provided, prompt if false
	},
}

// Parse as part of PreRunE
//var name string
//var start string
var length string
var end string

func init() {
 	rootCmd.AddCommand(startCmd)

	startCmd.Flags().StringVarP(&length, "length", "l", "5min", "How long the timer will run for.")

	current_time := time.Now()
	default_end := current_time.Add(time.Hour).String()
	startCmd.Flags().StringVarP(&end, "time", "t", default_end, "A specific time for the timer to end.")
}


// func parseStartInput(c *cobra.Command, args []string) error {
	
// }
