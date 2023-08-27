package ratelimiter

import (
	"github.com/SirNexus/go-tutorial/concurrency/internal/pkg/ratelimiter/bootstrap"
	"github.com/SirNexus/go-tutorial/concurrency/internal/pkg/ratelimiter/master"
	"github.com/spf13/cobra"
)

// Cmd runs our rate limiter test command
func Cmd() *cobra.Command {
	opts := &bootstrap.Options{}
	cmd := &cobra.Command{
		Use: "rate-limiter",
		RunE: func(cmd *cobra.Command, args []string) error {
			return master.NewMaster(opts).Run()
		},
	}

	opts.AddFlags(cmd)

	return cmd
}
