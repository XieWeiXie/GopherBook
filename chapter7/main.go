package main

import "github.com/spf13/cobra"

func init() {}

var ROOT = &cobra.Command{
	PreRun: func(cmd *cobra.Command, args []string) {

	},
	Run: func(cmd *cobra.Command, args []string) {

	},
	PostRun: func(cmd *cobra.Command, args []string) {

	},
}
