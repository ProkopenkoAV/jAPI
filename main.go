package main

import (
	"jAPI/cmd"
	"jAPI/cmd/create"
	"jAPI/cmd/del"
	"jAPI/cmd/running"
	"log"
)

func init() {
	log.Println("Initializing jAPI")
	registerCommands()
}

func registerCommands() {
	cmd.RootCmd.AddCommand(running.RunJobCmd)
	cmd.RootCmd.AddCommand(del.DelJobCmd)
	cmd.RootCmd.AddCommand(create.CreateJobCmd)
}

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
