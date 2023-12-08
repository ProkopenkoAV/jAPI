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
	RootCmd.PersistentFlags().StringP("server", "s", "", "Jenkins URL")
	RootCmd.PersistentFlags().StringP("port", "p", "", "Jenkins port")
	RootCmd.PersistentFlags().StringP("user", "u", "", "Jenkins user")
	RootCmd.PersistentFlags().StringP("token", "t", "", "Jenkins token")
	RootCmd.PersistentFlags().StringP("job", "j", "", "Job name")

	viper.BindPFlag("server", RootCmd.PersistentFlags().Lookup("server"))
	viper.BindPFlag("port", RootCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("user", RootCmd.PersistentFlags().Lookup("user"))
	viper.BindPFlag("token", RootCmd.PersistentFlags().Lookup("token"))
	viper.BindPFlag("job", RootCmd.PersistentFlags().Lookup("job"))
}
