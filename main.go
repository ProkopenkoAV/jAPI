package main

import (
	"fmt"
	"jAPI/cmd"
	"jAPI/cmd/create"
	"jAPI/cmd/delete"
	"jAPI/cmd/running"
	"jAPI/config"
)

func addCommand() {
	cmd.RootCmd.AddCommand(running.RunJobCmd)
	cmd.RootCmd.AddCommand(delete.DelJobCmd)
	cmd.RootCmd.AddCommand(create.CreateCmd)
}

func init() {
	fmt.Println("Initialization OK")
	addCommand()
}

func main() {
	config.InitConfig()
	cmd.Execute()
}
