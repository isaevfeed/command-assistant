package assistant

import (
	"encoding/json"
	"log"
	"os"
	"os/exec"

	"githab.com/command-assistant/internal/console"
	"githab.com/command-assistant/internal/helper"
)

type CommandAssistant struct {
	cons *console.Console
	data map[string]ProjectCommand
	cmd  Command
}

func Init() *CommandAssistant {
	return &CommandAssistant{
		&console.Console{},
		map[string]ProjectCommand{},
		Command{},
	}
}

func (ca *CommandAssistant) Start() error {
	ca.cons.Handle()
	file, err := ca.readConfig()
	defer file.Close()
	if err != nil {
		return err
	}

	err = json.NewDecoder(file).Decode(&ca.data)
	if err != nil {
		return err
	}

	if err := ca.Run(); err != nil {
		return err
	}

	return nil
}

func (ca *CommandAssistant) Run() error {
	for prj, cmd := range ca.data {
		if prj != ca.cons.Project {
			continue
		}

		cmdStruct, err := helper.GetFiledFromStruct(cmd, ca.cons.Command)
		if err != nil {
			return err
		}
		path, err := helper.GetStringFieldFromStruct(cmd, helper.FirstLettertoUpper(ca.cons.Stand))
		if err != nil {
			return err
		}

		if projCmd, ok := cmdStruct.(*Command); ok {
			projCmd.Args = append(projCmd.Args, path)
			cmd := exec.Command(projCmd.Command, projCmd.Args...)
			log.Printf("Commnad: %s \n", cmd)
			err = cmd.Run()
		}

		if _, ok := err.(*exec.ExitError); !ok {
			return err
		}

		log.Println("Done")
	}

	return nil
}

func (ca *CommandAssistant) readConfig() (*os.File, error) {
	file, err := os.Open(os.Getenv("COMMAND_CONFIG"))

	return file, err
}
