/*
 * @Time : 2019-06-25 17:52
 * @Author : ygqbasic@gmail.com
 * @File : cmd
 * @Software: VS Code
 */

package cmd

import (
	"flag"
	"github.com/spf13/cobra"
)

func AddFlags(rootCmd *cobra.Command) {
	flag.CommandLine.VisitAll(func(gf *flag.Flag) {
		rootCmd.PersistentFlags().AddGoFlag(gf)
	})
}
