package shapes

import strategy "github.com/max-prosper/go-design-patterns/behavioral/strategy/example2"

type TextSquare struct {
	strategy.DrawOutput
}

func (t *TextSquare) Draw() error {
	t.Writer.Write([]byte("Square"))
	return nil
}
