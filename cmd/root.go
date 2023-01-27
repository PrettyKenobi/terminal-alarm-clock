/*
Copyright © 2022 PrettyKenobi <prettykenobi@gmail.com>
*/
package cmd

import (
	"os"
	"runtime"
	"time"

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
	PersistentPreRun: checkForTimersFile,
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
	// when this action is called directly.s
}

func checkForTimersFile(cmd *cobra.Command, args []string) {
	logger, err := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()
	if err != nil {
		print("Could not make logger.")
	}

	currOs := runtime.GOOS

	switch currOs {
	case "windows":
		{
			// Save current directory
			originalDir := currFs.Name()
			// Checks for timers.json in user's directory
			err := currFs.Chdir("~")
			if err == nil {
				f, err := currFs.Stat("timers.json")
				if err == currFs.ErrNotExist {
					currFs.Create("timers.json")
					// TODO: Log that we created timers.json
					sugar.Infow("Created timers.json.",
						"time", time.Now(),
						"file", f.Name(),
					)
				}
				// Move back to original directory
				err = currFs.Chdir(currDir)
				if err != nil {
					sugar.Errorw("Couldn't go back to orginal directory", "time", time.Now(), "pwd", currFs.Name(), "original directory", currDir, "error", err)
				}
				sugar.Infow("Changed back to original directory",
					"time", time.Now(),
					"original directory", currDir)
			}
		}
	default:
		{
			// Check for timers.json & create if it doesn't exist
			f, err := currFs.Stat("timers.json")
			if err == currFs.Exists() {
				currFs.Create("timers.json")
				sugar.Infow("Created timers.json",
					"time", time.Now,
					"file_name", f.Name())
			}
		}
		//file exists
	}
}
