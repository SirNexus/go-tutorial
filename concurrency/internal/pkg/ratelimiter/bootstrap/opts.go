package bootstrap

import "github.com/spf13/cobra"

type Options struct {
	NumWorkers int
	RPS        float32
}

func (o *Options) AddFlags(cmd *cobra.Command) {
	cmd.Flags().IntVarP(&o.NumWorkers, "num-workers", "n", 2,
		"Number of workers to use when processing events")
	cmd.Flags().Float32VarP(&o.RPS, "requests-per-second", "r", 1,
		"Requests per second to allow through. Can be decimal")
}
