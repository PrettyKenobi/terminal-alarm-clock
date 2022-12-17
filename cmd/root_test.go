/*
Copyright © 2022 PrettyKenobi <prettykenobi@gmail.com>
*/

package cmd_test

import (
	"io/fs"
	"os"
	"runtime"
	"testing"
)

/* // Mock the OS

// Mock the filesystem
type mockedFS struct {
// Built in filesystem interface
fs.FS

reportName string
reportDir bool
}

// Mock FileInfo
type mockedFileInfo struct {
fs.FileInfo

// TODO: Add reportErr after making "checkForTimersFile" more robust wrt errors
name string
isDir bool
}

func (m mockedFileInfo) Name() string { return m.name }
func (m mockedFileInfo) IsDir() bool { return m.isDir } */


// Test "checkForTimersFile"
func TestcheckForTimersFile(t *testing.T) {
/* 	oldFS := os.DirFS("/")
	mfs := &mockedFS{}
	fs.FS = mfs
	// Change back to the real fs when done
	defer func ()  {
		fs = oldFS
	}() */

	// TODO: The actual testing
}

