package main

import (
	"fmt"
	"foo/bark"

	"github.com/spf13/cobra"
)

func main() {
	bark.Bark()
	fmt.Println(cobra.ShellCompRequestCmd)
}
