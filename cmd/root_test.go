/*
Copyright © 2022 PrettyKenobi <prettykenobi@gmail.com>
*/

package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

// Test that `checkForTimersFile` works correctly when the file doens't exist.
func TestCheckForTimersFile_NoFile(t *testing.T) {
	// Set up mock filesystem
	mockFs := afero.NewMemMapFs()
	// Save the real file system in order to switch back during cleanup
	realFs := currFs
	// Change to using mock filesystem
	currFs = mockFs
	// Create the user's home directory in the mock filesystem.
	mockUserHome, err := os.UserHomeDir()
	if err != nil {
		t.Fatal(err)
	}
	currFs.MkdirAll(mockUserHome, 0755)

	filePath := filepath.Join(mockUserHome, "tac-timers.json")
	exists, err := afero.Exists(currFs, filePath)
	if err != nil {
		t.Fatal(err)
	}
	if exists {
		t.Fatalf("Expected `tac-timers.json` to not exist but it does.")
	}

	rootCmd.Execute()

	exists, err = afero.Exists(currFs, filePath)
	assert.True(t, exists)

	// Reset currFs to the orginal filesystem
	defer func() { currFs = realFs }()
}

// Test that `checkForTimersFile` works correctly when the file already exists.
func TestCheckForTimersFile_Exists(t *testing.T) {
	// Set up mock filesystem
	mockFs := afero.NewMemMapFs()
	// Save the real file system in order to switch back during cleanup
	realFs := currFs
	// Change to using mock filesystem
	currFs = mockFs
	// Create the user's home directory in the mock filesystem.
	mockUserHome, err := os.UserHomeDir()
	if err != nil {
		t.Fatal(err)
	}
	currFs.MkdirAll(mockUserHome, 0755)

	filePath := filepath.Join(mockUserHome, "tac-timers.json")
	exists, err := afero.Exists(currFs, filePath)
	if err != nil {
		t.Fatal(err)
	}
	if !exists {
		t.Fatalf("Expected `tac-timers.json` to exist but it does not.")
	}

	rootCmd.Execute()

	exists, err = afero.Exists(currFs, filePath)
	assert.True(t, exists)

	// Reset currFs to the orginal filesystem
	defer func() { currFs = realFs }()
}
