package main

import (
	"github.com/jcyamacho/rxgen/cmd"
	"github.com/jcyamacho/rxgen/internal/funcutil"
)

func main() {
	funcutil.Fatal(cmd.Execute)()
}
