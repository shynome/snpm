package common

import (
	"fmt"
	"os"

	"github.com/shynome/snpm"
)

var Pkg snpm.Package

func init() {

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	Pkg, err = snpm.NewPackage(wd)
	if err != nil {
		fmt.Print(err, "\n")
		os.Exit(1)
	}

}
