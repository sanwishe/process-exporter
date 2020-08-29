package main

import (
	"fmt"
	"os"

	"github.com/sanwishe/processes-exporter/pkg/version"
	"github.com/spf13/cobra"
)

const (
	minArgsNum = 1
)

var (
	configFile string
	listenAddr string
)

func main() {
	root := &cobra.Command{
		Use:   "processes-exporter",
		Short: `Export detail metrics of processes`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < minArgsNum {
				return fmt.Errorf("processes_exporter need at least one argument: configfile")
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},

		Version: version.FullVersion(),
	}

	//root.Flags().AddFlag()

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
