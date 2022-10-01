/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a new alarm",
	Long: `Start a new alarm.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("start called")
		// TODO check if name provided, prompt for one if false
		// TODO check that name is unused, prompt for a different name if false
		// TODO check if duration provided, prompt if false
	},
}

var name string
var start int64
var end string

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	startCmd.Flags().StringP("name", "n", "The name of what you want to track.")
	startCmd.Flags().Var(&end, "duration", "d", "5min", "How long to set the alarm for.")

}
