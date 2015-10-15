package widgets

import (
	"fmt"
	"io"
	"os"

	"github.com/kildevaeld/go-acsii"
)

type WidgetConfig struct {
	MessageColor   acsii.Style
	HighlightColor acsii.Style
	StdinColor     acsii.Style
	Writer         io.Writer
}

var DefaultConfig WidgetConfig = WidgetConfig{
	MessageColor:   acsii.Dim,
	HighlightColor: acsii.Cyan,
	StdinColor:     acsii.Reset,
	Writer:         os.Stdout,
}

type Widget interface {
}

func write(w io.Writer, msg string, args ...interface{}) {
	w.Write([]byte(fmt.Sprintf(msg, args...)))
}
