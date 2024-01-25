package console

import (
	"log"
	"os"
)

type Console struct {
	Command string
	Project string
	Stand   string
}

func (c *Console) Handle() {
	c.checkArgs()

	c.Command = os.Args[1]
	c.Project = os.Args[2]
	c.Stand = os.Args[3]
}

func (c *Console) checkArgs() {
	if len(os.Args) < 4 {
		log.Fatal("Формат команды: ca {command:d|e} {project} {stand} {file_path}")
	}
}
