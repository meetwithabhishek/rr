//go:build fang

package cmd

import (
	"context"
	"os"

	"github.com/charmbracelet/fang"
)

func Execute() {
	if err := fang.Execute(context.Background(), rootCmd); err != nil {
		os.Exit(1)
	}
}
