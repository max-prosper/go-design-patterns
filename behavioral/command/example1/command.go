package main

import (
	"fmt"
)

type Command interface {
	Execute()
}

type ConsoleOutput struct {
	message string
}

func (c *ConsoleOutput) Execute() {
	fmt.Println(c.message)
}

func CreateCommand(s string) Command {
	fmt.Println("Creating command")

	return &ConsoleOutput{
		message: s,
	}
}

type CommandQueue struct {
	queue []Command
}

func (p *CommandQueue) AddCommand(c Command) {
	p.queue = append(p.queue, c)

	if len(p.queue) == 3 {
		for _, cmd := range p.queue {
			cmd.Execute()
		}

		p.queue = make([]Command, 3)
	}
}

func main() {
	queue := CommandQueue{}

	messages := []string{"First", "Second", "Third", "Fourth", "Fifth"}

	for i := 0; i < len(messages); i++ {
		queue.AddCommand(CreateCommand(messages[i] + " message"))
	}
}
