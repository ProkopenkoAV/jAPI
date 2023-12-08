package main

import (
	"fmt"
	"jAPI/cmd"
	"jAPI/cmd/create"
	"jAPI/cmd/delete"
	"jAPI/cmd/running"
)

func init() {
	fmt.Println("init...")
	registerCommands()
}

func registerCommands() {
	cmd.RootCmd.AddCommand(running.RunJobCmd)
	cmd.RootCmd.AddCommand(delete.DelJobCmd)
	cmd.RootCmd.AddCommand(create.CreateCmd)
}

func main() {
	cmd.Execute()
}
