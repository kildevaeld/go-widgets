package widgets

import (
	"fmt"
	"io"
	"os"

	"github.com/kildevaeld/go-ascii"
)

type WidgetConfig struct {
	MessageColor   ascii.Style
	HighlightColor ascii.Style
	StdinColor     ascii.Style
	Writer         io.Writer
}

var DefaultConfig WidgetConfig = WidgetConfig{
	MessageColor:   ascii.Dim,
	HighlightColor: ascii.Cyan,
	StdinColor:     ascii.Reset,
	Writer:         os.Stdout,
}

type Widget interface {
}

func write(w io.Writer, msg string, args ...interface{}) int {
	str := fmt.Sprintf(msg, args...)
	w.Write([]byte(str))
	return len(str)
}
