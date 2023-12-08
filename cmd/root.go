package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var RootCmd = &cobra.Command{
	Use:   "Jenkins-cli",
	Short: "Jenkins cli client",
	Long:  "A command line utility of Jenkins",
}

func init() {
	RootCmd.PersistentFlags().StringP("server", "s", "", "Jenkins URL")
	RootCmd.PersistentFlags().StringP("port", "p", "", "Jenkins port")
	RootCmd.PersistentFlags().StringP("user", "u", "", "Jenkins user")
	RootCmd.PersistentFlags().StringP("token", "t", "", "Jenkins token")
	RootCmd.PersistentFlags().StringP("job", "j", "", "Job name")

	if err := viper.BindPFlag("server", RootCmd.PersistentFlags().Lookup("server")); err != nil {
		log.Fatalf("Failed to bind flag: %v", err)
	}
	if err := viper.BindPFlag("port", RootCmd.PersistentFlags().Lookup("port")); err != nil {
		log.Fatalf("Failed to bind flag: %v", err)
	}
	if err := viper.BindPFlag("user", RootCmd.PersistentFlags().Lookup("user")); err != nil {
		log.Fatalf("Failed to bind flag: %v", err)
	}
	if err := viper.BindPFlag("token", RootCmd.PersistentFlags().Lookup("token")); err != nil {
		log.Fatalf("Failed to bind flag: %v", err)
	}
	if err := viper.BindPFlag("job", RootCmd.PersistentFlags().Lookup("job")); err != nil {
		log.Fatalf("Failed to bind flag: %v", err)
	}
}
