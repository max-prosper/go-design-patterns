package shapes

import (
	"fmt"
	"os"

	"github.com/max-prosper/go-design-patterns/behavioral/strategy/example2"
)

const (
	TextStrategy  = "text"
	ImageStrategy = "image"
)

func Factory(s string) (strategy.Output, error) {
	switch s {
	case TextStrategy:
		return &TextSquare{
			DrawOutput: strategy.DrawOutput{
				LogWriter: os.Stdout,
			},
		}, nil
	case ImageStrategy:
		return &ImageSquare{
			DrawOutput: strategy.DrawOutput{
				LogWriter: os.Stdout,
			},
		}, nil
	default:
		return nil, fmt.Errorf("Strategy '%s' not found", s)
	}
}
