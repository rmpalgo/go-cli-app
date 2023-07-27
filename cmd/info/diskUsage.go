package info

import (
	"fmt"

	"github.com/ricochet2200/go-disk-usage/du"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var diskUsageCmd = &cobra.Command{
	Use:   "diskUsage",
	Short: "Print disk usage of the current directory",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		defaultDirectory := "."

		if dir := viper.GetViper().GetString("cmd.info.diskUsage.defaultDirectory"); dir != "" {
			defaultDirectory = dir
		}

		fmt.Printf(
			"THE DISK USAGE DIRECTORY BEING USED: %s\n",
			defaultDirectory,
		)

		usage := du.NewDiskUsage(defaultDirectory)

		fmt.Printf(
			"Free disk space: %+v in directory %s\n",
			usage.Used(),
			defaultDirectory,
		)
	},
}

func init() {
	InfoCmd.AddCommand(diskUsageCmd)
}
