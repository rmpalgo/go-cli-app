package net

import (
	"github.com/spf13/cobra"
)

var NetCmd = &cobra.Command{
	Use:   "net",
	Short: "Net contains network based commands",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {}
