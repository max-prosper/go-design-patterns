package chain

import (
	"fmt"
	"io"
	"strings"
)

type ChainLogger interface {
	Next(string)
}

type FirstLogger struct {
	NextChain ChainLogger
}

func (fl *FirstLogger) Next(s string) {
	fmt.Printf("First logger: %s\n", s)

	if fl.NextChain != nil {
		fl.NextChain.Next(s)
	}
}

type SecondLogger struct {
	NextChain ChainLogger
}

func (sl *SecondLogger) Next(s string) {
	if strings.Contains(strings.ToLower(s), "hello") {
		fmt.Printf("Second logger: %s\n", s)

		if sl.NextChain != nil {
			sl.NextChain.Next(s)
		}

		return
	}

	fmt.Printf("Finishing in second logger\n\n")
}

type WriterLogger struct {
	NextChain ChainLogger
	Writer    io.Writer
}

func (wl *WriterLogger) Next(s string) {
	if wl.Writer != nil {
		wl.Writer.Write([]byte("WriterLogge: " + s))
	}

	if wl.NextChain != nil {
		wl.NextChain.Next(s)
	}
}

type ClosureChain struct {
	NextChain ChainLogger
	Closure   func(string)
}

func (c *ClosureChain) Next(s string) {
	if c.Closure != nil {
		c.Closure(s)
	}

	if c.NextChain != nil {
		c.Next(s)
	}
}
