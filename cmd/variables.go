package cmd

import (
	"github.com/spf13/afero"
)

var currFs = afero.NewOsFs()
var prod = false

//var muFS sync.Mutex

// var muProd sync.Mutex
