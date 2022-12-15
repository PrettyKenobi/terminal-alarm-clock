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
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
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
				// Checks for timers.json in user's directory
				err := os.Chdir("~")
				if err == nil {
					f, err := os.Stat("timers.json")
					if err == os.ErrNotExist {
						os.Create("timers.json")
						// TODO: Log that we created timers.json
						sugar.Infow("Created timers.json.",
							"time", time.Now(),
							"file", f.Name(),
						)
					}
				}
			}
		default:
			{
				// Check for .config/terminal-alarm-clock, create if it doesn't exist
				err := os.Chdir("~/.config")
				if err == nil {
					err := os.Chdir("terminal-alarm-clock")
					if err == os.ErrNotExist {
						err := os.Mkdir("terminal-alarm-clock", 0750)
						if err == nil {
							// TODO: Log that we made the directory
							sugar.Infow("Created tac folder",
								"time", time.Now(),
								"dir_name", "terminal-alarm-clock",
							)
							os.Chdir("terminal-alarm-clock")
						}
					}
					// Our config directory exists so time to check for timers.json
					f, err := os.Stat("timers.json")
					if err == os.ErrNotExist {
						os.Create("timers.json")
						sugar.Infow("Created timers.json",
							"time", time.Now,
							"file_name", f.Name())
					}
				}
			}
			//file exists
		}
	},
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
