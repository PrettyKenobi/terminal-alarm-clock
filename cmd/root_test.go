/*
Copyright © 2022 PrettyKenobi <prettykenobi@gmail.com>
*/

package cmd

import (
	"testing"

	"github.com/spf13/afero"
)

func TestCheckForTimersFile_NoFile(t *testing.T) {
	// Set up mock filesystem
	mockFs := afero.NewMemMapFs()
	// Save the real file system in order to switch back during cleanup
	realFs := currFs
	// Change to using mock filesystem
	currFs = mockFs

	if !currFs.Exists("~/timers.json") {

	}
	rootCmd.Execute()

}
