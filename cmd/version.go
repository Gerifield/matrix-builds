package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/the-gophers/gophers/xcobra"
)

var (
	// GitCommit is the git reference injected at build
	GitCommit string
)

func newVersionCommand() (*cobra.Command, error) {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the git ref",
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) error {
			fmt.Println(GitCommit)
			return nil
		}),
	}, nil
}
