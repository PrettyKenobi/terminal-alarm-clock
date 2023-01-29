/*
Copyright © 2022 PrettyKenobi <prettykenobi@gmail.com>
*/
package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// Rootcmd Represents The Base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "terminal-alarm-clock",
	Short: "An easy to use egg-timer that runs on the terminal",
	Long:  `TODO`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	// PersistentPreRun: checkForTimersFile,
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

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//checkForTimersFile()
}

// See if a file called `tac-timers.json` exists in the user's home directory.
// Create the file if it doesn't exist.
func checkForTimersFile() {
	// Create logger with zap
	// Create development logger in development environment
	var logger, err = zap.NewDevelopment()
	if prod {
		// Create production logger if in production
		logger, err = zap.NewProduction()
	}
	defer logger.Sync()
	// Use the more relaxed version of the zap logger
	sugar := logger.Sugar()
	if err != nil {
		print("Could not make logger.")
	}

	// Get user's home diretory.
	homeDir, err := os.UserHomeDir()
	if err != nil {
		sugar.Errorw("Cannot get user home directory")
	}
	// Check for `tac-timers.json` in the home directory
	filePath := filepath.Join(homeDir, ".tac-timers.json")
	exists, err := afero.Exists(currFs, filePath)
	if err != nil {
		sugar.Errorw("Unable to check for `tac-timers.json`")
	}
	if !exists {
		// File doesn't exist
		currFs.Create(filePath)
		sugar.Infow("Created `tac-timers.json`.")
	} else {
		//File does exist.
		sugar.Infow("`tac-timers.json` already exists.")
	}
}
