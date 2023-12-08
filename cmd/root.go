package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCmd = &cobra.Command{
	Use:   "Jenkins-cli",
	Short: "Jenkins cli client",
	Long:  "A command line utility of Jenkins",
}

func Execute() {
	cobra.CheckErr(RootCmd.Execute())
}

func init() {
	RootCmd.PersistentFlags().StringP("url", "url", "", "Jenkins URL")
	RootCmd.PersistentFlags().StringP("port", "port", "", "Jenkins port")
	RootCmd.PersistentFlags().StringP("user", "user", "", "Jenkins user")
	RootCmd.PersistentFlags().StringP("token", "token", "", "Jenkins token")

	viper.BindPFlag("url", RootCmd.PersistentFlags().Lookup("url"))
	viper.BindPFlag("port", RootCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("user", RootCmd.PersistentFlags().Lookup("user"))
	viper.BindPFlag("token", RootCmd.PersistentFlags().Lookup("token"))

}
