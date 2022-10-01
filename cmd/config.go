/*
Copyright © 2022 PrettyKenobi <prettykenobi@gmail.com>

*/
package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "terminal-alarm-clock config [flags]",
	Short: "View and edit configuration settings",
	Long: `View and edit configuration settings for terminal-alarm-clock`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	configCmd.PersistentFlags().String()
}

