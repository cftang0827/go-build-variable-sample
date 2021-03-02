package main

import (
	"fmt"

	"go-build-variable-sample/types"
)

func main() {
	fmt.Printf("Hello world, version: %s\n", types.GitCommit)
}
